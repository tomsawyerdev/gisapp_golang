package models

import (
	//"database/sql"
	"context"
	"encoding/json"
	//"errors"
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
		  as zonif FROM zonif h WHERE  h.fieldid=$1 ORDER BY creation `

	rows, err := svc.DB.Query(context.Background(), sql, id) //FieldId

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

	//fmt.Println("Rows count:", len(records))
	//fmt.Println("Record:", records[0])

	//for idxRow, row := range records {			zonif := row.Zonif}
	// Color Assignation
	// --------------------------------

	for idxRow, row := range records {

		vra := row.Zonif

		//fmt.Println("Count:", vra.ZoneCount)
		//fmt.Println("Base Colors:", vra.Colors)
		//fmt.Printf("t1: %T\n", vra.Colors)

		pallete := grd.HexList2RgbList(vra.Colors)
		// Genero los colores
		var colors []string
		if vra.ZoneCount == 1 { // OJO Da error si count==1
			colors = []string{"#008000"} // verde
		} else {
			colors = grd.HexMultiColorScale(pallete, vra.ZoneCount)
		}
		//fmt.Println("colors:", colors)
		// Asigno los colores a  cada poligono de la zona

		//https://blog.davidvassallo.me/2022/05/11/golang-gotcha-modifying-an-array-of-structs/
		for idxZone, zz := range vra.Zones {

			//fmt.Println("zone:", idxZone, zz.Name)
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
	const sql = `SELECT zonif_boundary_buffer(@id,@name,@distance);`

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

// ----------------------------------------------
func ZoneCreate(body dto.ZoneCreate) (map[string]string, error) {

	//fmt.Println("ZoneCreate")
	var sql string
	sql = `SELECT * FROM boundary_intersection(@zonifid,@jsonpol)` // call store procedure, return text:Error or WKTPolygon

	//fmt.Printf("t1: %T\n", body.Polygon) // map[string]interface{}
	//fmt.Println(body.Polygon)

	jsonString, _ := json.Marshal(body.Polygon)

	args := pgx.NamedArgs{"zonifid": body.ZonifId, "jsonpol": jsonString}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	var result string
	err := row.Scan(&result) //return: Error or WKTPolygon

	//fmt.Println("Scan boundary_intersection:", result, err)
	//fmt.Println("result:", result["boundary_intersection"][0:5])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return nil, err
	}

	if result[0:5] == "Error" {

		return map[string]string{"status": "error", "msg": result[6:]}, nil
	}

	//return map[string]string{"status": "error", "msg": "resulterror"}, nil

	// 2) Corto con el boundary
	//fmt.Println("Clip:", body.Clip)
	args = pgx.NamedArgs{"zonifid": body.ZonifId, "wktpol": result}
	if body.Clip == "o" {
		// Difference sometimes return polygon so use ST_Multi(
		// ST_Multi(ST_Difference(polygons,ST_PolygonFromText(%(wktpol)s)))
		// ST_Difference can return GEOMETRYCOLLECTION EMPTY
		sql = `UPDATE zonifitems SET polygons= (WITH dif AS( select ST_Multi(ST_SimplifyVW(ST_Difference(st_setsrid(polygons,4326),st_setsrid(ST_PolygonFromText(@wktpol),4326)),0.0000001)) as geo)
				SELECT CASE WHEN ST_IsEmpty(geo)  THEN 'MULTIPOLYGON EMPTY'::geometry ELSE geo END FROM dif )
				WHERE zonifid=@zonifid`

		_, err = svc.DB.Exec(context.Background(), sql, args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
			return nil, err
		}

		// Delete if some is empty
		sql = `DELETE FROM zonifitems WHERE zonifid=@zonifid AND ST_isEmpty(polygons) `
		_, err = svc.DB.Exec(context.Background(), sql, args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
			return nil, err
		}

		sql = `INSERT into zonifitems(zonifid,name,editable,clip,zoneorder,polygons)
		VALUES (@zonifid,@name,TRUE,'o',
		(SELECT max(zoneorder)+1 FROM zonifitems WHERE zonifid=@zonifid),
		st_multi(ST_CollectionExtract(st_makevalid(ST_PolygonFromText(@wktpol)),3)))`
		// st_makevalid cuando hay error devuelve GEOMETRY COLLECTION

		args = pgx.NamedArgs{"zonifid": body.ZonifId, "name": body.Name, "wktpol": result}

		_, err = svc.DB.Exec(context.Background(), sql, args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
			return nil, err
		}

	} else {
		// Clip Under

		sql = `SELECT  st_AsText(ST_Difference(st_setsrid(ST_GeomFromText(@wktpol),4326),
		(SELECT  ST_Union(st_setsrid(polygons,4326)) FROM zonifitems WHERE zonifid=@zonifid GROUP BY zonifid)
		 )) as geometry`

		rows, err := svc.DB.Query(context.Background(), sql, args)

		if err != nil {
			fmt.Println("Rows err:", err)
			return nil, err
		}
		defer rows.Close()

		records, err := pgx.CollectRows(rows, pgx.RowToMap)
		if err != nil {
			fmt.Println("Rows err:", err)
			return nil, err
		}

		//---------------------------------

		wktdifpol := records[0]["geometry"].(string)
		//fmt.Println("Difference pol:", wktdifpol[0:7])

		if wktdifpol[0:7] == "POLYGON" { // SINGLE POLYGON

			sql = `INSERT into zonifitems(zonifid,name,editable,clip,zoneorder,polygons)
			VALUES (@zonifid,@name,TRUE,'u',
			(SELECT max(zoneorder)+1 FROM zonifitems WHERE zonifid=@zonifid),
			 ST_Multi(ST_setsrid(ST_PolygonFromText(@wktdifpol),4326)))`

			args = pgx.NamedArgs{"zonifid": body.ZonifId, "name": body.Name, "wktdifpol": wktdifpol}

			_, err = svc.DB.Exec(context.Background(), sql, args)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
				return nil, err
			}

		} else {
			//MULTIPLYGON
			sql = `SELECT ST_astext(dump.poly) as wktonepoly FROM (SELECT (ST_Dump(ST_MultiPolygonFromText(@wktdifpol))).geom as poly ) AS dump  ORDER BY ST_Area(dump.poly) DESC LIMIT 1;`
			args = pgx.NamedArgs{"wktdifpol": wktdifpol}

			row := svc.DB.QueryRow(context.Background(), sql, args)

			var result map[string]any

			err := row.Scan(&result)
			//fmt.Println("Scan:", result)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
				return nil, err
			}

			wktonepoly := result["wktonepoly"].(string)
			//fmt.Println("wktonepoly:", wktonepoly[0:7])

			sql = `INSERT into zonifitems(zonifid,name,editable,clip,zoneorder,polygons)
			 VALUES (@zonifid,@name,TRUE,'u',
			 (SELECT max(zoneorder)+1 FROM zonifitems WHERE zonifid=@zonifid),
			  ST_Multi(ST_setsrid(ST_PolygonFromText(@wktonepoly),4326)))`

			args = pgx.NamedArgs{"zonifid": body.ZonifId, "name": body.Name, "wktonepoly": wktonepoly}

			_, err = svc.DB.Exec(context.Background(), sql, args)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
				return nil, err
			}

		}

	}
	return nil, nil
}

// High Geometry Complexity
func ZoneUpdBoundary(body dto.ZoneUpdBoundary) (map[string]string, error) {

	fmt.Printf("ZoneUpdBoundary")

	var sql string
	sql = `SELECT * FROM boundary_intersection(@zonifid,@jsonpol)` // call store procedure, return text:Error or WKTPolygon

	fmt.Printf("t1: %T\n", body.Polygon) // map[string]interface{}
	//fmt.Println(body.Polygon)

	jsonString, _ := json.Marshal(body.Polygon)

	args := pgx.NamedArgs{"zonifid": body.ZonifId, "jsonpol": jsonString}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	var result string
	err := row.Scan(&result) //return: Error or WKTPolygon

	fmt.Println("Scan boundary_intersection:", result, err)
	//fmt.Println("result:", result["boundary_intersection"][0:5])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return nil, err
	}

	if result[0:5] == "Error" {

		return map[string]string{"status": "error", "msg": result[6:]}, nil
	}

	//return map[string]string{"status": "error", "msg": "resulterror"}, nil

	// 2) Corto con el boundary
	fmt.Println("Clip:", body.Clip)
	args = pgx.NamedArgs{"zonifid": body.ZonifId, "zoneid": body.ZoneId, "wktpol": result}
	if body.Clip == "o" {

		// Difference sometimes return polygon so use ST_Multi(
		// ST_Multi(ST_Difference(polygons,ST_PolygonFromText(%(wktpol)s)))
		// ST_Difference can return GEOMETRYCOLLECTION EMPTY
		sql = `UPDATE zonifitems SET polygons= (WITH dif AS( select ST_Multi(ST_Difference(st_setsrid(polygons,4326),st_setsrid(ST_PolygonFromText(@wktpol),4326))) as geo)
											SELECT CASE WHEN ST_IsEmpty(geo)  THEN 'MULTIPOLYGON EMPTY'::geometry ELSE geo END FROM dif                               
											)
		WHERE zonifid=@zonifid AND id != @zoneid`

		_, err = svc.DB.Exec(context.Background(), sql, args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
			return nil, err
		}

		// Delete if some is empty
		sql = `DELETE FROM zonifitems WHERE zonifid=@zonifid AND ST_isEmpty(polygons) `
		_, err = svc.DB.Exec(context.Background(), sql, args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
			return nil, err
		}

		// Update the zone polygon
		sql = `UPDATE zonifitems SET polygons= ST_Multi(ST_PolygonFromText(@wktpol))
			WHERE zonifid=@zonifid AND id=@zoneid`

		_, err = svc.DB.Exec(context.Background(), sql, args)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
			return nil, err
		}

		return nil, nil

	} else {
		// Clip is Under
		sql = `SELECT  st_AsText(ST_Difference(st_setsrid(ST_GeomFromText(@wktpol),4326),
				(SELECT  ST_Union(st_setsrid(polygons,4326))  FROM zonifitems  WHERE zonifid=@zonifid  AND id!=@zoneid GROUP BY zonifid)
	 			)) as geometry`

		rows, err := svc.DB.Query(context.Background(), sql, args)

		if err != nil {
			fmt.Println("Rows err:", err)
			return nil, err
		}
		defer rows.Close()

		records, err := pgx.CollectRows(rows, pgx.RowToMap)
		if err != nil {
			fmt.Println("Rows err:", err)
			return nil, err
		}

		//---------------------------------

		wktdifpol := records[0]["geometry"].(string)
		fmt.Println("difference pol:", wktdifpol[0:7])

		if wktdifpol[0:7] == "POLYGON" { // SINGLE POLYGON

			sql = `UPDATE zonifitems SET polygons= ST_Multi(ST_SimplifyVW(ST_PolygonFromText(@wktdifpol),0.0000001))
				WHERE zonifid=@zonifid AND id=@zoneid`

			args := pgx.NamedArgs{"zonifid": body.ZonifId, "zoneid": body.ZoneId, "wktdifpol": wktdifpol}
			_, err = svc.DB.Exec(context.Background(), sql, args)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
				return nil, err
			}

		} else {
			//MULTIPOLYGON then Extract max area polygon

			fmt.Println("MULTIPOLYGON")

			sql = `SELECT ST_astext(dump.poly) as wktonepoly FROM (SELECT (ST_Dump(ST_MultiPolygonFromText(@wktdifpol))).geom as poly ) AS dump  ORDER BY ST_Area(dump.poly) DESC LIMIT 1;`

			args := pgx.NamedArgs{"wktdifpol": wktdifpol}
			row := svc.DB.QueryRow(context.Background(), sql, args)

			var result map[string]any

			err := row.Scan(&result)
			fmt.Println("Scan:", result)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
				return nil, err
			}

			wktonepoly := result["wktonepoly"].(string)
			fmt.Println("wktonepoly:", wktonepoly[0:7])

			sql = `UPDATE zonifitems SET polygons=  ST_Multi(ST_SimplifyVW(ST_PolygonFromText(@wktonepoly),0.0000001))                
				WHERE zonifid=@zonifid AND id=@zoneid`

			args = pgx.NamedArgs{"zonifid": body.ZonifId, "zoneid": body.ZoneId, "wktonepoly": wktonepoly}
			_, err = svc.DB.Exec(context.Background(), sql, args)

			if err != nil {
				fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
				return nil, err
			}

		}
	}
	return nil, nil
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
	const sql = `UPDATE zonifitems set clip=@clip WHERE id=@id AND zonifid=@zonifid ;`
	args := pgx.NamedArgs{"id": body.ZoneId, "zonifid": body.ZonifId, "clip": clip}
	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}
