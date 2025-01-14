package dao // data access objects

// EL `db:"XXX" no hace falta
type ZonifBody struct {
	Id        string      `db:"id" json:"id"` // Definido como ::text en el SQL
	FieldId   string      `db:"fieldid" json:"fieldid"`
	Name      string      `db:"name" json:"name"`
	Source    string      `json:"source"`
	Colors    []string    `json:"colors"`
	Creation  interface{} `json:"creation"` //creation:2022-01-01
	ZoneCount int         `json:"zonecount" `
	Zones     []Zone      `json:"zones"`
}

type ZonifKey struct {
	Zonif ZonifBody `json:"zonif" `
}
