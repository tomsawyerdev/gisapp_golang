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
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ZonifUpdColors struct {
	Id     int      `json:"id" binding:"required"`
	Colors []string `json:"colors" binding:"required"`
}
type ZonifDelete struct {
	Id   int    `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
}
type ZonifCreateBuffer struct {
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	Distance int    `json:"distance" binding:"required"`
}
type ZoneRename struct {
	Id      int    `json:"zoneid" binding:"required"`
	ZonifId int    `json:"zonifid" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

type ZoneDelete struct {
	Id      int    `json:"id" binding:"required"`
	ZonifId int    `json:"zonifid" binding:"required"`
	Name    string `json:"name" binding:"required"`
}

// ----------------------------------------------------
type ZoneCreate struct {
	ZonifId int         `json:"zonifid" binding:"required"`
	Name    string      `json:"name" binding:"required"`
	Polygon interface{} `json:"polygon" binding:"required"`
	Clip    string      `json:"clip" binding:"required"`
}
type ZoneUpdBoundary struct {
	ZoneId  int         `json:"zoneid" binding:"required"`
	ZonifId int         `json:"zonifid" binding:"required"`
	Polygon interface{} `json:"polygon" binding:"required"`
	//Clip    string `json:"clip" binding:"required"`
}

//----------------------------------------------------

type ZoneRemovePoints struct {
	ZoneId  int `json:"zoneid" binding:"required"`
	ZonifId int `json:"zonifid" binding:"required"`
}
type ZoneSimplify struct {
	ZoneId  int `json:"zoneid" binding:"required"`
	ZonifId int `json:"zonifid" binding:"required"`
}

type ZoneRefine struct {
	ZoneId  int `json:"zoneid" binding:"required"`
	ZonifId int `json:"zonifid" binding:"required"`
}
type ZoneUpdClip struct {
	ZoneId  int    `json:"zoneid" binding:"required"`
	ZonifId int    `json:"zonifid" binding:"required"`
	Clip    string `json:"clip" binding:"required"`
}
