package models

import (
	//"database/sql"
	"context"
	//"encoding/json"
	"fmt"
	"gisapi/dao"
	svc "gisapi/database"
	"gisapi/dto"
	"github.com/jackc/pgx/v5"
	//"log"
	grd "gisapi/colors"
	"os"
)

// *sql.Rows
//type Record []byte //map[string]any

//{"id" : 11, "name" : "La Candelaria", "centroid" : {"type":"Point","coordinates":[-63.68448692,-31.321594989]}}

// Vra Sql
const sql = `SELECT  json_build_object(  
		'id', v.id::text,    
		'fieldid', v.fieldid::text,        
		'zonifid', v.zonifid::text,        
		'zonifname', (select name  FROM zonif WHERE id = v.zonifid ),    
		'creation', v.creation,            
		'target', v.target,         
		'name', v.name,       
		'obs', v.obs,    
		'zonecount',(select count(*)  FROM zonifitems zit WHERE zit.zonifid= v.zonifid ),            
		'colors',(select colors FROM zonif  WHERE id= v.zonifid ),            
		'zones', ARRAY(SELECT row_to_json(r)
						 FROM (
						 SELECT it.id,it.zonifid,it.name,
								to_char(it.minvalue,'FM9990.000') as minvalue,
								to_char(it.maxvalue,'FM9990.000') as maxvalue,
								to_char(ST_Area(polygons::geography)/10000, 'FM999999999.00') as area,
								ARRAY (SELECT  json_build_object(
												  'type','Feature',
												  'geometry',(st_AsGeoJSON((ST_DUMP(polygons)).geom))::json ) ) polygons 
												   FROM zonifitems it WHERE it.zonifid= v.zonifid ORDER BY it.minvalue) r ),
												   
		'channels', ARRAY(SELECT row_to_json(r)
						 FROM (
						 SELECT ch.id,ch.name,ch.unit,ch.values
							   FROM vrachannels ch WHERE ch.vraid= v.id ORDER BY ch.id) r )                                                            
		  )
		  as vra FROM vra v WHERE fieldid = $1 ORDER BY creation`

//fmt.Println("ListFarms----------")
//fmt.Println("Database Ping:", svc.DB.Ping(context.Background()) == nil)

//args := pgx.NamedArgs{"fieldid": fieldid}

func VraList(fieldid int) ([]dao.VraKey, error) {

	rows, err := svc.DB.Query(context.Background(), sql, fieldid)

	if err != nil {
		fmt.Println("Rows err:", err)
		return nil, err
	}
	defer rows.Close()

	//records, err := pgx.CollectRows(rows, pgx.RowToMap)
	records, err := pgx.CollectRows(rows, pgx.RowToStructByName[dao.VraKey])

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return nil, err
	}

	fmt.Println("Rows count:", len(records))
	//fmt.Println("Record:", records[0])

	for idxRow, row := range records {

		vra := row.Vra

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
			records[idxRow].Vra.Zones[idxZone].Color = colors[idxZone]
			for idxPol, _ := range zz.Polygons { // # asigno a cada poligono de la zona
				//pol.Color = "#008000" //colors[idx]
				//p2 := &pol
				//p2.SetColor("#008000")
				records[idxRow].Vra.Zones[idxZone].Polygons[idxPol].Color = colors[idxZone]

			}

		}

	}

	return records, nil

}

func VraList2(fieldid int) ([]map[string]any, error) {

	rows, err := svc.DB.Query(context.Background(), sql, fieldid)

	if err != nil {
		fmt.Println("Rows err:", err)
		return nil, err
	}
	defer rows.Close()

	records, err := pgx.CollectRows(rows, pgx.RowToMap)
	//records, err := pgx.CollectRows(rows, pgx.RowToStructByName[dao.vra]

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return nil, err
	}

	fmt.Println("Rows count:", len(records))
	//fmt.Println("Record:", records[0])

	//COLOR ASSIGNATION
	// Tiene que ser ordenada :menor a mayor ORDER BY it.minvalue o por zone

	for _, r := range records {
		var colors []string
		vra := r["vra"].(map[string]any)
		zonecount := int(vra["zonecount"].(float64))

		fmt.Println("Count:", zonecount)
		basecolors := vra["colors"] // es []interface {} hay que castearlo to	.([]string)
		fmt.Println("Base Colors:", basecolors)
		fmt.Printf("t1: %T\n", basecolors)
		break

		pallete := grd.HexList2RgbList(basecolors.([]string))
		//count=r['vra']
		//['zonecount']

		// https://www.reddit.com/r/golang/comments/bsd6bc/converting_an_interface_to_string/
		// UNmarshall to struct ???

		// Genero los colores
		if zonecount == 1 { // OJO Da error si count==1
			colors = []string{"#008000"} // verde
		} else {
			colors = grd.HexMultiColorScale(pallete, zonecount)
		}
		fmt.Println("colors:", colors)
		// Asigno los colores a  cada poligono de la zona
		zones := vra["zones"].([]map[string]any)
		for idx, zz := range zones {

			fmt.Println("zones:", idx, zz)

		}

	}

	/*
	   for r in rows:
	       pallete=grd.HexList2RgbList(r['vra']['colors'])
	       count=r['vra']['zonecount']
	       #print(grad)
	       if (count>1): # OJO Da error si count=1
	           colors=grd.HexMultiColorScale(pallete,count)
	       else:
	           colors=['#008000'] # verde
	       #print(colors)

	       for idx,zz in  enumerate(r['vra']['zones']):
	          #print(z)
	          zz['color']=colors[idx] # a cada zona
	          for pol in zz['polygons']: # a cada poligono de la zona
	              pol['color']=colors[idx]


	*/

	return records, nil

}

// Hacer una transaccion

func VraCreate(body dto.VraCreate) error {

	const sql = `INSERT into vra(zonifid,fieldid,creation,target,name,obs) 
	VALUES( @zonifid,@fieldid,CURRENT_DATE,@target,@name,@obs) 
	RETURNING id`

	args := pgx.NamedArgs{"zonifid": body.ZonifId, "fieldid": body.FieldId, "name": body.Name, "obs": body.Obs}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	//Scan the returning id

	fmt.Println("Row:", row)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed INSERT rows: %v\n", err)
		return err
	}

	const sql2 = `INSERT into vrachannels(vraid,name,unit,values) VALUES( @vraid,@name,@unit,@values) `

	//Warning: executemany()

	tags, err = svc.DB.Exec(context.Background(), sql2, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed INSERT rows: %v\n", err)
		return err
	}

	return nil

}

func VraRename(body dto.VraRename) error {

	const sql = `UPDATE vra SET name=@name WHERE id=@id ;`
	args := pgx.NamedArgs{"id": body.Id, "name": body.Name}
	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed UPDATE rows: %v\n", err)
		return err
	}

	return nil
}

func VraDelete(body dto.VraDelete) error {
	return nil
}

//-----------------------------------------------
// Channel
//-----------------------------------------------

func VraChannelCreate(body dto.VraChannel) error {

	const sql = `INSERT into vrachannels(vraid,name,unit,values) VALUES( @vraid,@name,@unit,@values)`
	args := pgx.NamedArgs{"vraid": body.VraId, "name": body.Name, "unit": body.Unit, "values": body.Values}
	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return err
	}

	return nil
}

func VraChannelRename(body dto.VraChannelRename) error {

	const sql = `UPDATE vrachannels SET name=@name WHERE id=@channelid AND vraid=@vraid ;`
	args := pgx.NamedArgs{"channelid": body.ChannelId, "vraid": body.VraId, "name": body.Name}
	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return err
	}

	return nil
}

func VraChannelUpdate(body dto.VraChannelUpdate) error {

	const sql = `UPDATE vrachannels SET values=@values,unit=@unit WHERE id=@id AND vraid=@vraid ;`
	args := pgx.NamedArgs{"id": body.Id, "vraid": body.VraId, "name": body.Name, "unit": body.Unit, "values": body.Values}
	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return err
	}

	return nil
}

func VraChannelDelete(body dto.VraChannelDelete) error {
	const sql = `DELETE from vrachannels  WHERE id=@channelid AND vraid=@vraid ;`
	args := pgx.NamedArgs{"channelid": body.ChannelId, "vraid": body.VraId}
	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed db action: %v\n", err)
		return err
	}

	return nil
}
