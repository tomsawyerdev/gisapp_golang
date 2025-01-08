package models

import (
	//"database/sql"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	svc "gisapi/database"
	"gisapi/dto"
	"github.com/jackc/pgx/v5"
	"os"
)

// -----------------------------------------
// Fields
// -----------------------------------------

func FieldsList(body dto.FieldList) ([]map[string]any, error) {

	const sql = `SELECT row_to_json(row) as field
                  FROM (select 'Feature' as type, fields.id as id, fields.name as name,
                  to_char(ST_Area(boundary::geography)/10000, 'FM999999999.00') as area,
                  st_AsGeoJSON(st_centroid(boundary))::json as centroid,
                  st_AsGeoJSON(boundary)::json as geometry,
                  coalesce(iscircle,False)::text as iscircle,
                  st_x(center) as centerlng,
                  st_y(center) as centerlat,
                  radius
                  FROM farms inner join fields on farms.id = fields.farmid where userid = @userid and farms.id= @farmid ORDER BY name) row;`

	args := pgx.NamedArgs{"userid": body.UserId, "farmid": body.FarmId}
	rows, err := svc.DB.Query(context.Background(), sql, args)

	//fmt.Println("Rows err:", err)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	records, err := pgx.CollectRows(rows, pgx.RowToMap)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed collecting rows: %v\n", err)
	}

	fmt.Println("Rows count:", len(records))
	//fmt.Println("Record:", records[0])
	return records, nil

}

func FieldVerifyOwnership(body dto.FieldCreate) error {

	const sql = `SELECT count(farms.id) as count FROM users, farms WHERE users.id=@userid 
				AND farms.userid = users.id          
				AND farms.id=@farmid;`

	args := pgx.NamedArgs{"userid": body.UserId, "farmid": body.FarmId}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	//fmt.Println("Rows err:", err)
	// Row count
	var count int

	err := row.Scan(&count)
	if err != nil {
		//fmt.Println("err:", err)
		return err
	}
	//fmt.Println("Rows count:", count)
	if count != 1 {
		return errors.New("FieldVerifyOwnership")
	}
	return nil
}

func FieldcreateCircle(body dto.FieldCreate) error {

	const sql = `INSERT into fields(farmid,name,iscircle,center,radius,boundary)
            VALUES( @farmid,
                    @name,
                    TRUE,
                    ST_MakePoint(@lng,@lat),
                    CAST(@radius as integer),
                    ST_Buffer(ST_MakePoint(@lng,@lat)::geography, @radius)::geometry);`

	args := pgx.NamedArgs{"farmid": body.FarmId, "name": body.Name, "lng": body.Lng, "lat": body.Lat, "radius": body.Radius}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

// map[string]any
func FieldcreateVerifyBoundary(geojson string) (bool, string, error) {

	//sql :=`With  t(g) as (select ST_GeomFromGeoJSON(@geojson)))
	//        Select ST_isValid(g) isvalid, ST_IsValidReason(g) reason from t;`

	const sql = `With  t(g) as (select ST_GeomFromGeoJSON(@geojson))
          Select  json_build_object('isvalid',ST_isValid(g),
                                    'reason',ST_IsValidReason(g)) as verify from t;`

	args := pgx.NamedArgs{"geojson": geojson}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	var verify map[string]any
	//fmt.Println("Pre Scan:")
	err := row.Scan(&verify) // verify ṕuede ser nil

	//fmt.Println("Verify Boundary:", verify)

	if err != nil {
		//fmt.Println("err:", err)
		return false, "", err
	}

	return verify["isvalid"].(bool), verify["reason"].(string), nil
}

func FieldcreatePolygon(body dto.FieldCreate) error {

	const sql = `INSERT into fields(farmid,name,iscircle,center,radius,boundary)
            VALUES( @farmid,
                    @name,
                    FALSE,
                    null,
                    null,
                    ST_GeomFromGeoJSON(@polygon));`

	geojson, _ := json.Marshal(body.Polygon)
	args := pgx.NamedArgs{"farmid": body.FarmId, "name": body.Name, "polygon": geojson}

	_, err := svc.DB.Exec(context.Background(), sql, args)
	//rows, err := svc.DB.Query(context.Background(), sql, args)
	//fmt.Println("FieldcreatePolygon rows:", rows)
	//fmt.Println("FieldcreatePolygon rows:", tags.RowsAffected())
	//fmt.Println("FieldcreatePolygon err:", err)

	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func Fieldrename(body dto.FieldRename) error {

	//const sql  = `UPDATE fields SET name= @name) WHERE farmid::text=@farmid) AND id::text=@fieldid);"
	const sql = `UPDATE fields SET name=@name WHERE farmid=@farmid::integer AND id=@fieldid::integer;`

	args := pgx.NamedArgs{"farmid": body.FarmId, "fieldid": body.FieldId, "name": body.Name}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func Fielddelete(body dto.FieldDelete) error {

	const sql = `DELETE FROM fields WHERE id=@fieldid AND farmid=@farmid AND trim(name)=trim(@name);`

	args := pgx.NamedArgs{"farmid": body.FarmId, "fieldid": body.FieldId, "name": body.Name}

	_, err := svc.DB.Exec(context.Background(), sql, args)

	//fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		//log.Fatal(err)
		return err
	}

	return nil
}

func FieldboundaryCircle(body dto.FieldBoundary) error {
	const sql = `UPDATE fields SET iscircle=TRUE,
                         center=ST_MakePoint(@lng,@lat),
                         radius =  CAST(@radius as integer),
                         boundary = ST_Buffer(ST_MakePoint(@lng,@lat)::geography, CAST(@radius as float) )::geometry
                         WHERE farmid=@farmid AND id=@fieldid;`

	args := pgx.NamedArgs{"farmid": body.FarmId, "fieldid": body.FieldId, "lng": body.Lng, "lat": body.Lat, "radius": body.Radius}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		//log.Fatal(err)
		return err
	}

	return nil
}

func FieldboundaryPolygon(body dto.FieldBoundary) error {
	const sql = `UPDATE fields SET iscircle=FALSE,
                        center=null,
                        radius = null,
                        boundary = ST_GeomFromGeoJSON(@polygon)
                        WHERE farmid=@farmid AND id=@fieldid;`

	geojson, _ := json.Marshal(body.Polygon)
	args := pgx.NamedArgs{"farmid": body.FarmId, "fieldid": body.FieldId, "polygon": geojson}

	_, err := svc.DB.Exec(context.Background(), sql, args)

	if err != nil {
		//log.Fatal(err)
		return err
	}
	return nil
}

func Fielddownloadgeojson(id int) (string, error) {

	const sql = `select concat(name,'.json') as name,json_build_object(
    'type', 'FeatureCollection',
    'features', array(select json_build_object(
           'type', 'Feature',
           'geometry',st_AsGeoJSON(boundary)::json,
           'properties',json_build_object('id',id::text,
                                          'name',name,
                                          'area',concat(to_char(ST_Area(boundary::geography)/10000, 'FM999999999.00'),'ha')
                                          )
            )))
            as fields from fields where id=@id;`

	args := pgx.NamedArgs{"id": id}

	row := svc.DB.QueryRow(context.Background(), sql, args)

	var geojson map[string]any
	var name string

	err := row.Scan(&name, &geojson) // verify ṕuede ser nil

	//fmt.Println("Verify Boundary:", verify)

	if err != nil {
		//fmt.Println("err:", err)
		return "", err
	}

	geostr, _ := json.Marshal(geojson)
	//fmt.Println("Row Scan:", name, string(geostr))
	return string(geostr), nil
}
