package models

import (
	//"database/sql"
	"context"
	//"encoding/json"
	"errors"
	"fmt"
	"gisapi/dao"
	svc "gisapi/database"
	"gisapi/dto"

	grd "gisapi/colors"

	"os"

	"github.com/jackc/pgx/v5"
)

// -----------------------------------------
// Zonifications
// -----------------------------------------

func ZonifList(id int) ([]dao.ZonifKey, error) {

	const sql = `SELECT  json_build_object(  
		'id', h.id::text,         
		'fieldid', h.fieldid::text,
		'name', h.name,   
		'source', h.source,             
		'creation', h.creation,   
		'colors', h.colors,                      
		'zonecount',(select count(*)  FROM zonifitems it WHERE it.zonifid= h.id ),            
		'zones', ARRAY(SELECT row_to_json(r)
						 FROM (
						 SELECT it.id::text,it.zonifid::text,it.name,it.zoneorder::text,
								coalesce(it.editable,false) as editable,
								clip,
								true as visible,
								to_char(it.minvalue,'FM9990.000') as minvalue,
								to_char(it.maxvalue,'FM9990.000') as maxvalue,
								to_char(ST_Area(polygons::geography)/10000, 'FM999999999.00') as area,
								ST_NumGeometries(polygons) as polycount,
								ARRAY (SELECT  json_build_object(
												  'type','Feature',
												  'geometry',(st_AsGeoJSON((ST_DUMP(polygons)).geom))::json ) ) polygons 
												   FROM zonifitems it WHERE it.zonifid=h.id ORDER BY it.zoneorder) r )        
		 )
		  as zonif FROM zonif h WHERE  h.fieldid=@fieldid ORDER BY creation `

	rows, err := svc.DB.Query(context.Background(), sql, id)

	if err != nil {
		fmt.Println("Rows err:", err)
		return nil, err
	}
	defer rows.Close()

	//records, err := pgx.CollectRows(rows, pgx.RowToMap)
	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[dao.ZonifKey])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return nil, err
	}

	fmt.Println("Rows count:", len(records))
	//fmt.Println("Record:", records[0])

	//for idxRow, row := range records {			zonif := row.Zonif}
	// Color Assignation
	// --------------------------------

	for idxRow, row := range records {

		vra := row.Zonif

		fmt.Println("Count:", vra.ZoneCount)
		fmt.Println("Base Colors:", vra.Colors)
		//fmt.Printf("t1: %T\n", vra.Colors)

		pallete := grd.HexList2RgbList(vra.Colors)
		// Genero los colores
		var colors []string
		if vra.ZoneCount == 1 { // OJO Da error si count==1
			colors = []string{"#008000"} // verde
		} else {
			colors = grd.HexMultiColorScale(pallete, vra.ZoneCount)
		}
		fmt.Println("colors:", colors)
		// Asigno los colores a  cada poligono de la zona

		//https://blog.davidvassallo.me/2022/05/11/golang-gotcha-modifying-an-array-of-structs/
		for idxZone, zz := range vra.Zones {

			fmt.Println("zone:", idxZone, zz.Name)
			//zz.Color = "#008000"
			//p1 := &zz                         //colors[idx]            // asigno  a cada zona
			//p1.SetColor("#008000")            //colors[idx]            // asigno  a cada zona
			records[idxRow].Zonif.Zones[idxZone].Color = colors[idxZone]
			for idxPol := range zz.Polygons { // # asigno a cada poligono de la zona
				//pol.Color = "#008000" //colors[idx]
				//p2 := &pol
				//p2.SetColor("#008000")
				records[idxRow].Zonif.Zones[idxZone].Polygons[idxPol].Color = colors[idxZone]

			}

		}

	}

	return records, nil

}

func ZonifCreate(body dto.ZonifCreate) error {

	const sql = `select zonif_create(@fieldid,@name);` // Call store procedure, insert zonif and zonif_items
	args := pgx.NamedArgs{"fieldid": body.FieldId, "name": body.Name}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

func ZonifRename(body dto.ZonifRename) error {
	const sql = `UPDATE zonif SET name=@name WHERE id=@id ;`
	args := pgx.NamedArgs{"id": body.Id, "name": body.Name}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
func ZonifUpdColors(body dto.ZonifUpdColors) error {
	const sql = `UPDATE zonif SET colors=@colors WHERE id=@id ;`
	args := pgx.NamedArgs{"id": body.Id, "colors": body.Colors}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
func ZonifDelete(body dto.ZonifDelete) error {

	// ON DELETE CASCADE;
	const sql = `DELETE FROM zonif WHERE id=@id AND name=@name;`
	args := pgx.NamedArgs{"id": body.Id, "name": body.Name}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

func ZonifCreateBuffer(body dto.ZonifCreateBuffer) error {
	const sql = `select zonif_boundary_buffer(@id,@name,@distance);`

	args := pgx.NamedArgs{"id": body.Id, "name": body.Name, "distance": body.Distance}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

// -----------------------------------------------------------
// Zone
//-----------------------------------------------------------

func ZoneRename(body dto.ZoneRename) error {

	const sql = `UPDATE zonifitems SET name=@name WHERE id=@id AND zonifid=@zonifid;`

	args := pgx.NamedArgs{"id": body.Id, "zonifid": body.ZonifId, "name": body.Name}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
func ZoneDelete(body dto.ZoneDelete) error {

	const sql = `DELETE FROM zonifitems  WHERE id=@id AND zonifid=@zonifid AND name=@name;`
	args := pgx.NamedArgs{"id": body.Id, "zonifid": body.ZonifId, "name": body.Name}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

// High Geometry Complexity
func ZoneCreate(body dto.ZoneCreate) (map[string]string, error) {

	sql := `SELECT * FROM boundary_intersection(%(zonifid)s,%(wktpol)s)` // call ctore procedure

	fmt.Printf("t1: %T\n", body.Polygon)
	fmt.Println(body.Polygon)

	args := pgx.NamedArgs{"zonifid": body.ZonifId, "wktpol": body.Polygon}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	var result map[string]string

	err := row.Scan(&result)
	fmt.Println("Scan:", result, err)

	if result["boundary_intersection"][0:5] == "Error" {

		//return map[string]string{"status": "error","msg":result["boundary_intersection"][6:]},nil
		return map[string]string{"status": "error", "msg": result["boundary_intersection"][6:]}, nil

	}

	//wktpol := result

	fmt.Println("clip:", body.Clip)

	err = errors.New("not implementet yet")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return nil, err
	}

	return nil, nil
}

// High Geometry Complexity
func ZoneUpdBoundary(body dto.ZoneUpdBoundary) error {

	err := errors.New("not implementet yet")

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
func ZoneRemovePoints(body dto.ZoneRemovePoints) error {

	const sql = `UPDATE zonifitems set polygons=ST_RemoveRepeatedPoints(polygons,0.0005) WHERE id=@id AND zonifid=@zonifid ;`
	args := pgx.NamedArgs{"id": body.ZoneId, "zonifid": body.ZonifId}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
func ZoneSimplify(body dto.ZoneSimplify) error {
	const sql = `UPDATE zonifitems set polygons=ST_SimplifyVW(polygons,0.000007) WHERE id=@id AND zonifid=@zonifid ;`
	args := pgx.NamedArgs{"id": body.ZoneId, "zonifid": body.ZonifId}
	_, err := svc.DB.Exec(context.Background(), sql, args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
func ZoneRefine(body dto.ZoneRefine) error {
	const sql = `UPDATE zonifitems set polygons=(
		SELECT st_multi(dump.poly) FROM (SELECT (ST_Dump(polygons)).geom as poly ) AS dump  ORDER BY ST_Area(dump.poly) DESC LIMIT 1
		) WHERE id=@id AND zonifid=@zonifid ;`

	args := pgx.NamedArgs{"id": body.ZoneId, "zonifid": body.ZonifId}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

func ZoneUpdClip(body dto.ZoneUpdClip) error {
	// Flip cLip value

	clip := "u"
	if body.Clip == "u" {
		clip = "o"
	}
	const sql = `sql="""UPDATE zonifitems set clip=@clip WHERE id=@id AND zonifid=@zonifid ;`
	args := pgx.NamedArgs{"id": body.ZoneId, "zonifid": body.ZonifId, "clip": clip}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
