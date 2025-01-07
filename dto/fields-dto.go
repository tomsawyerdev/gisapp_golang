package dto // data transfer objects

//Warning: si uso el mismo dto para todos no le puedo poner el required

type FieldList struct {
	UserId int
	FarmId int `json:"farmid" binding:"required"`
}

type FieldCreate struct {
	UserId  int
	FarmId  int            `json:"farmid"`
	Name    string         `json:"name" binding:"required,max=100"`
	Type    string         `json:"type" binding:"required,max=20"`
	Lat     float32        `json:"lat"`
	Lng     float32        `json:"lng"`
	Radius  float32        `json:"radius"`
	Polygon map[string]any `json:"polygon"`
}

type FieldRename struct {
	UserId  int
	FarmId  int    `json:"farmid" binding:"required"`
	FieldId int    `json:"fieldid" binding:"required"`
	Name    string `json:"name" binding:"required,max=255"`
}

type FieldDelete FieldRename

// Update Boundary

type FieldBoundary struct {
	UserId  int
	FarmId  int            `json:"farmid"`
	FieldId int            `json:"fieldid" binding:"required"`
	Type    string         `json:"type" binding:"required,max=20"`
	Lat     float32        `json:"lat"`
	Lng     float32        `json:"lng"`
	Radius  float32        `json:"radius"`
	Polygon map[string]any `json:"polygon"`
}
