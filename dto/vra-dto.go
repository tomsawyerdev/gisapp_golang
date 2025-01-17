package dto // data transfer objects

//Warning si uso el mismo dto para todos no le puedo poner el required

type VraList struct {
	FieldId int `json:"fieldid" binding:"required"`
}

type VraCreate struct {
	ZonifId  int    `json:"zonifid" binding:"required"`
	FieldId  int    `json:"fieldid" binding:"required"`
	Name     string `json:"name" binding:"required,max=255" sqlParameterName:"name"`
	Obs      string `json:"obs" binding:"required,max=255" `
	Channels []VraChannel
}

// se reutiliza arriba
type VraChannel struct {
	VraId  int       `json:"vraid"` // binding:"required"
	Name   string    `json:"name" binding:"required,max=255"`
	Unit   string    `json:"unit" binding:"required,max=20" `
	Values []float32 `json:"values" `
}

type VraRename struct {
	Id   string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required,max=255"`
}

type VraDelete VraRename

type VraChannelCreate struct {
	//Id     int       `json:"id"` // binding:"required"
	VraId  string    `json:"vraid"` // binding:"required"
	Name   string    `json:"name" binding:"required,max=255"`
	Unit   string    `json:"unit" binding:"required,max=20" `
	Values []float32 `json:"values" `
}

type VraChannelRename struct {
	ChannelId int    `json:"channelId" binding:"required"`
	VraId     string `json:"vraid" binding:"required"`
	Name      string `json:"name" binding:"required"`
}
type VraChannelUpdate struct {
	Id     int       `json:"id" binding:"required"` // Warning
	VraId  string    `json:"vraid" binding:"required"`
	Name   string    `json:"name" binding:"required"`
	Unit   string    `json:"unit" binding:"required,max=20" `
	Values []float32 `json:"values" `
}
type VraChannelDelete struct {
	ChannelId int    `json:"channelid" binding:"required"`
	Name      string `json:"name" binding:"required"`
	VraId     string `json:"vraid" binding:"required"`
}
