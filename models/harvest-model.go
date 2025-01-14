package models

import (
	"context"
	//"encoding/json"
	"fmt"
	svc "gisapi/database"
	"gisapi/dto"
	"github.com/jackc/pgx/v5"
	//"log"

	"os"
)

func HarvestList(body dto.HarvestList) ([]map[string]any, error) {
	const sql = `select json_build_object(
                'id', hs.id,
                'fieldid',hs.fieldid,
                'crop', hs.crop,
                'daystart', hs.daystart,
                'dayend', hs.dayend,
                'operations', ARRAY(SELECT row_to_json(r)
                            FROM (select id,machine,operator,capture::text,(select count(*) from harveststamps where harveststamps.harvestoperationid=ho.id ) pointcount FROM harvestoperations ho where hs.id = ho.seasonid ORDER BY capture) r))
            as season FROM harvestseason hs where fieldid=@fieldid ORDER BY daystart DESC;`

	//CAST($(fieldid) as integer)

	args := pgx.NamedArgs{"fieldid": body.FieldId}

	rows, err := svc.DB.Query(context.Background(), sql, args)

	//fmt.Println("Rows err:", err)
	if err != nil {

		return nil, err
	}
	defer rows.Close()

	records, err := pgx.CollectRows(rows, pgx.RowToMap)

	if err != nil {
		fmt.Fprintf(os.Stderr, "failed collecting rows: %v\n", err)
	}

	//fmt.Println("Rows count:", count)
	//fmt.Println("Record:", records[0])
	return records, nil

}

func HarvestSeasonCreate(body dto.HarvestSeasonCreate) error {
	const sql = `INSERT into  harvestseason(fieldid,crop,daystart,dayend)
              values(@fieldid,@name,CURRENT_DATE,CURRENT_DATE + INTERVAL '1 month');`

	args := pgx.NamedArgs{"fieldid": body.FieldId, "name": body.Name}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

func HarvestSeasonUpdate(body dto.HarvestSeasonUpdate) error {
	const sql = ` UPDATE harvestseason SET
                        crop=@name,
                        daystart=to_date(@daystart, 'YYYY/MM/DD'),
                        dayend=to_date(@dayend, 'YYYY/MM/DD')
                        WHERE id=@id;`

	args := pgx.NamedArgs{"id": body.Id, "name": body.Name, "daystart": body.DayStart, "dayend": body.DayEnd}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

func HarvestSeasonDelete(body dto.HarvestSeasonDelete) error {
	const sql = `DELETE FROM harvestseason WHERE id=@id;`
	args := pgx.NamedArgs{"id": body.Id}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}
	return nil
}

//----------------------------------------------------------------------
// Operations
// ----------------------------------------------------------------------

func HarvestOperationCreate(body dto.HarvestOperationCreate) error {
	const sql = `INSERT into harvestoperations(
                            seasonid,
                            operator,
                            machine,
                            capture)
                        VALUES( @seasonid,
                                @name,
                                @machine,
                                to_date(@capture, 'YYYY/MM/DD')
                                );`
	args := pgx.NamedArgs{"seasonid": body.SeasonId, "name": body.Name, "machine": body.Machine, "capture": body.Capture}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}

	return nil
}

func HarvestOperationUpdate(body dto.HarvestOperationUpdate) error {
	const sql = `UPDATE harvestoperations SET
                operator=@name,
                machine=@machine,
                capture=to_date(@capture, 'YYYY/MM/DD')
                WHERE id= @id;`
	args := pgx.NamedArgs{"id": body.Id, "name": body.Name, "machine": body.Machine, "capture": body.Capture}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}
	return nil

}

func HarvestOperationDelete(body dto.HarvestOperationDelete) error {
	const sql = `DELETE FROM harvestoperations WHERE id=@id;`

	args := pgx.NamedArgs{"id": body.Id}

	tags, err := svc.DB.Exec(context.Background(), sql, args)

	fmt.Println("tags rows:", tags.RowsAffected())

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return err
	}
	return nil
}

//https://stackoverflow.com/questions/36233566/inserting-multiple-records-with-pg-promise
//https://stackoverflow.com/questions/37300997/multi-row-insert-with-pg-promise

/*
func Harvestimportlogcsv(body map[string]any) error {
	const sql = ` `
	return nil
}*/

// ----------------------------------------------------------------------
// Operations Histogram
// ----------------------------------------------------------------------
func HarvestOperationsValues(body dto.HarvestOperationsHist) ([]float32, error) {

	var columns = []string{"yield", "speed", "moisture", "headwidth", "headheight", "altitude"}

	varIdx := body.Variable
	col := columns[0]

	if varIdx >= 0 && varIdx < 6 {
		col = columns[varIdx]
	}

	//fmt.Println("column:", col)

	sql := "SELECT array_agg(COALESCE( " + col + ",0.0)) as values FROM harveststamps WHERE harvestoperationid = ANY(@hoids);"

	args := pgx.NamedArgs{"hoids": body.Hoids}
	row := svc.DB.QueryRow(context.Background(), sql, args)

	var values []float32
	//fmt.Println("Pre Scan:")
	err := row.Scan(&values) // verify ṕuede ser nil

	//fmt.Println("Verify Boundary:", verify)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return []float32{}, err
	}

	return values, nil

}

// Return: [{x,y,value}]
func HarvestOperationStamps(colIdx int, hoids []int) ([]dto.HarvestOperationsStamps, error) {

	//fmt.Println("HarvestOperationStamps-----------------")

	var columns = []string{"yield", "speed", "moisture", "headwidth", "headheight", "altitude"}

	varIdx := colIdx
	col := columns[0]

	if varIdx >= 0 && varIdx < 6 {
		col = columns[varIdx]
	}

	var sql = `WITH s1 AS ( SELECT st_transform(st_setsrid(location,4326),3857) g, ` + col + ` as value
    from harveststamps WHERE harvestoperationid = ANY(@hoids) )
    select st_x(g) x,st_y(g) y, COALESCE(value,0) as value, ARRAY[0,0,255] as color from s1;`

	args := pgx.NamedArgs{"hoids": hoids}

	rows, err := svc.DB.Query(context.Background(), sql, args)

	//fmt.Println("Rows err:", err)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return nil, err
	}
	defer rows.Close()

	//stamps, err := pgx.CollectRows(rows, pgx.RowToMap)
	//stamps, err := pgx.CollectRows(rows, pgx.RowToStructByName[dto.HarvestOperationsStamps]) // No acepta Variables de mas
	stamps, err := pgx.CollectRows(rows, pgx.RowToStructByPos[dto.HarvestOperationsStamps]) // No acepta Variables de mas

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: failed collecting rows: %v\n", err)
		return nil, nil
	}

	//fmt.Println("Rows count:", rows.CommandTag().RowsAffected())
	//fmt.Println("Record:", stamps[0])
	return stamps, nil

}

// Image Bounding Box 4326
// se llama con un POST
// devuelve un json
// se envia [[-31.394199557,-63.700761988],[-31.391267406,-63.69500004]]
func HarvestOperationsBounds4326(hoids []int) ([][]float32, error) {
	// body = %{"hoids" = [2,3]} ??
	//IO.puts("HarvestModel.harvestoperationBounds4326")
	//IO.inspect(body, label: "hoids")

	const sql = `WITH s1 AS ( SELECT st_extent(location) g from harveststamps WHERE harvestoperationid  = ANY(@hoids) group by  harvestoperationid  )
    SELECT row_to_json(bbox) bbox FROM (select min(st_xmin(g)) lonmin, min(st_ymin(g)) latmin, max(st_xmax(g)) lonmax, max(st_ymax(g)) latmax from s1) as bbox;`

	args := pgx.NamedArgs{"hoids": hoids}

	row := svc.DB.QueryRow(context.Background(), sql, args)

	var bbox map[string]float32
	//fmt.Println("Pre Scan:")
	err := row.Scan(&bbox)

	latmin := bbox["latmin"]
	lonmin := bbox["lonmin"]
	latmax := bbox["latmax"]
	lonmax := bbox["lonmax"]

	values := [][]float32{{latmin, lonmin}, {latmax, lonmax}}

	//bounds = [[ b['latmin'],b['lonmin']],[b['latmax'],b['lonmax']]]

	//fmt.Printf("Corner min: %f, %f, \n", latmin, lonmin)
	//fmt.Printf("corner max: %f, %f, \n", latmax, lonmax)

	//values :=[][]float32{{0,0},{10,10}}

	//fmt.Println("Bounds 4326:", bbox)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		//return [][]float32{}, err
		return nil, err
	}

	return values, nil

}

// Image Bounding Box 3857
// llamado interno

func HarvestOperationsBounds3857(hoids []int) (map[string]float64, error) {

	const sql = `WITH s1 AS ( SELECT st_extent(st_transform(st_setsrid(location,4326),3857)) g
                   from harveststamps WHERE harvestoperationid = ANY(@hoids) group by  harvestoperationid  )
                   SELECT row_to_json(bbox) bbox FROM
                   (select min(st_xmin(g)) xmin,min(st_ymin(g)) ymin, max(st_xmax(g)) xmax, max(st_ymax(g)) ymax from s1) as bbox`

	args := pgx.NamedArgs{"hoids": hoids}

	row := svc.DB.QueryRow(context.Background(), sql, args)

	var bbox map[string]float64
	//fmt.Println("Pre Scan:")
	err := row.Scan(&bbox)

	//fmt.Println("Bounds 3857:", bbox)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Fail db operation : %v\n", err)
		return map[string]float64{}, err
	}

	//fmt.Printf("t1: %T\n", bbox["xmin"])

	return bbox, nil
}
