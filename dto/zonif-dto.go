package dto // data transfer objects

//Warning si uso el mismo dto para todos no le puedo poner el required

type ZonifList struct {
	FieldId int `json:"fieldid" binding:"required"`
}

type ZonifCreate struct {
	FieldId int    `json:"fieldid" binding:"required"`
	Name    string `json:"name" binding:"required"`
}
type ZonifRename struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ZonifUpdColors struct {
	Id     string   `json:"id" binding:"required"`
	Colors []string `json:"colors" binding:"required"`
}
type ZonifDelete struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ZonifCreateBuffer struct {
	Id       string `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Distance string `json:"distance" binding:"required"`
}
type ZoneRename struct {
	Id      string `json:"zoneid" binding:"required"`
	ZonifId string `json:"zonifid" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type ZoneDelete struct {
	Id      int    `json:"id" binding:"required"`
	ZonifId string `json:"zonifid" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

// ----------------------------------------------------
type ZoneCreate struct {
	ZonifId string      `json:"zonifid" binding:"required"`
	Name    string      `json:"name" binding:"required"`
	Polygon interface{} `json:"polygon" binding:"required"`
	Clip    string      `json:"clip" binding:"required"`
}
type ZoneUpdBoundary struct {
	ZoneId  string      `json:"zoneid" binding:"required"`
	ZonifId string      `json:"zonifid" binding:"required"`
	Polygon interface{} `json:"polygon" binding:"required"`
	Clip    string      `json:"clip" binding:"required"`
	//Type  string      `json:"type" binding:"required"`
}

//----------------------------------------------------

type ZoneRemovePoints struct {
	ZoneId  string `json:"zoneid" binding:"required"`
	ZonifId string `json:"zonifid" binding:"required"`
}
type ZoneSimplify struct {
	ZoneId  string `json:"zoneid" binding:"required"`
	ZonifId string `json:"zonifid" binding:"required"`
}

type ZoneRefine struct {
	ZoneId  string `json:"zoneid" binding:"required"`
	ZonifId string `json:"zonifid" binding:"required"`
}
type ZoneUpdClip struct {
	ZoneId  string `json:"zoneid" binding:"required"`
	ZonifId string `json:"zonifid" binding:"required"`
	Clip    string `json:"clip" binding:"required"`
}
