package dto // data transfer objects

//Warning si uso el mismo dto para todos no le puedo poner el required

type FarmNew struct {
	Id     int    `sqlParameterName:"id"`
	Userid int    `sqlParameterName:"userid"`
	Name   string `json:"name" binding:"required,max=255" sqlParameterName:"name"`
}
