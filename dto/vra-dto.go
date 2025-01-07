package dto // data transfer objects

//Warning si uso el mismo dto para todos no le puedo poner el required

type VraCreate struct {
	ZonifId int    `json:"zonifid" binding:"required"`
	FieldId int    `json:"fieldid" binding:"required"`
	Name    string `json:"name" binding:"required,max=255" sqlParameterName:"name"`
	Obs     string `json:"obs" binding:"required,max=255" `
}

type VraChannel struct {
	VraId  int       `json:"vraid" binding:"required"`
	Name   string    `json:"name" binding:"required,max=255"`
	Unit   string    `json:"units" binding:"required,max=20" `
	Values []float32 `json:"values" `
}

type VraRename struct {
	Id   int    `json:"vraid" binding:"required"`
	Name string `json:"name" binding:"required,max=255"`
}

type VraDelete VraRename

type VraChannelCreate VraChannel

type VraChannelRename struct {
	ChannelId int    `json:"channelId" binding:"required"`
	VraId     int    `json:"vraid" binding:"required"`
	Name      string `json:"name" binding:"required"`
}
type VraChannelUpdate struct {
	Id     int       `json:"id" binding:"required"` // Warning
	VraId  int       `json:"vraid" binding:"required"`
	Name   string    `json:"name" binding:"required"`
	Unit   string    `json:"units" binding:"required,max=20" `
	Values []float32 `json:"values" `
}
type VraChannelDelete struct {
	ChannelId int `json:"channelid" binding:"required"`
	VraId     int `json:"vraid" binding:"required"`
}
