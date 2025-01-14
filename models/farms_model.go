package models

import (
	//"database/sql"
	"context"
	//"encoding/json"
	"fmt"
	svc "gisapi/database"
	"gisapi/dto"
	"github.com/jackc/pgx/v5"
	//"log"
	"os"
)

// *sql.Rows
//type Record []byte //map[string]any

//{"id" : 11, "name" : "La Candelaria", "centroid" : {"type":"Point","coordinates":[-63.68448692,-31.321594989]}}

func FarmsList(userid int) ([]map[string]any, error) {

	const sql = `select  json_build_object(
	  'id', farms.id,
	  'name', farms.name,
	  'centroid', (SELECT st_AsGeoJSON(st_centroid(ST_Extent(boundary))) FROM fields where farms.id = fields.farmid )::json
	  )
	  as farm FROM farms where userid = $1 ORDER BY name;`

	//fmt.Println("ListFarms----------")
	//fmt.Println("Database Ping:", svc.DB.Ping(context.Background()) == nil)

	rows, err := svc.DB.Query(context.Background(), sql, userid)

	//fmt.Println("Rows err:", err)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return nil, err
	}
	defer rows.Close()

	//http://go-database-sql.org/retrieving.html
	//https://go.dev/doc/tutorial/database-access#multiple_rows
	//count := 0
	//fmt.Println("Rows count:", count)
	// Print ColumnsTypes
	/*
		fields := rows.FieldDescriptions()
		for _, v := range fields {

			fmt.Println("Fields:", v.Name, v.Format)
		} */

	//	Extract Rows to Json
	//var records []map[string]any

	/*
		for rows.Next() {
			count += 1

			var raw []byte
			err := rows.Scan(&raw)
			//fmt.Println("rows.Scan, err:", err)
			if err != nil {
				return nil, err
			}
			var result map[string]any
			json.Unmarshal(raw, &result)

			records = append(records, result)
		}*/

	records, err := pgx.CollectRows(rows, pgx.RowToMap)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed collecting rows: %v\n", err)
	}

	//fmt.Println("Rows count:", count)
	//fmt.Println("Record:", records[0])
	return records, nil

}

func FarmCreate(body dto.FarmNew) error {

	const sql = `INSERT into farms(userid,name) VALUES(@userid,@name) ;`

	args := pgx.NamedArgs{"userid": body.Userid, "name": body.Name}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil

}

func FarmUpdate(body dto.FarmNew) error {

	const sql = `UPDATE farms SET name=@name WHERE userid=@userid AND id=@id;`

	args := pgx.NamedArgs{"userid": body.Userid, "id": body.Id, "name": body.Name}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}
	return nil
}

func FarmDelete(body dto.FarmNew) error {

	const sql = `DELETE FROM farms WHERE userid=@userid AND id=@id AND name=@name;`

	args := pgx.NamedArgs{"userid": body.Userid, "id": body.Id, "name": body.Name}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())
	fmt.Println("tags rows:", tags.Delete())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil

}
