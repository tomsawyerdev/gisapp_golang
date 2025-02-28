package dao // data access objects

//zones:[map[area:5.36 id:10 maxvalue:<nil> :<nil> :L2 polygons:

// EL `db:"XXX" no hace falta
type VraBody struct {
	Id        string        `db:"id" json:"id"` // Definido como ::text en el SQL
	ZonifId   string        `db:"zonifid" json:"zonifid"`
	ZonifName string        `db:"zonifname" json:"zonifname"`
	FieldId   string        `db:"fieldid" json:"fieldid"`
	Name      string        `db:"name" json:"name"`
	Obs       string        `db:"obs" json:"obs"`
	Channels  []interface{} `json:"channels"` //id:3 name:N1 unit:sem/ha values:[10 20 30 40 50]
	Colors    []string      `json:"colors"`
	Creation  interface{}   `json:"creation"` //creation:2022-01-01
	Target    string        `json:"target"`
	ZoneCount int           `json:"zonecount" `
	Zones     []Zone        `json:"zones"`
}

type VraKey struct {
	Vra VraBody `json:"vra" `
}
