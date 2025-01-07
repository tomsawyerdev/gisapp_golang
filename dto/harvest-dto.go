package dto // data transfer objects

type HarvestList struct {
	UserId  int
	FieldId int `json:"fieldid" binding:"required"`
}

type HarvestSeasonCreate struct {
	UserId  int
	FieldId int    `json:"fieldid" binding:"required"`
	Name    string `json:"name" binding:"required,max=100"`
}

type HarvestSeasonUpdate struct {
	UserId   int
	Id       int    `json:"id" binding:"required"`
	Name     string `json:"name" binding:"required,max=100"`
	DayStart string `json:"daystart" binding:"required,max=10"`
	DayEnd   string `json:"dayend" binding:"required,max=10"`
}

type HarvestSeasonDelete struct {
	UserId int
	Id     int `json:"id" binding:"required"`
}

type HarvestOperationCreate struct {
	UserId   int
	SeasonId int    `json:"seasonid" binding:"required"`
	Name     string `json:"name" binding:"required,max=100"`
	Machine  string `json:"machine" binding:"required,max=100"`
	Capture  string `json:"capture" binding:"required,max=10"`
}

type HarvestOperationUpdate struct {
	UserId  int
	Id      int    `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required,max=100"`
	Machine string `json:"machine" binding:"required,max=100"`
	Capture string `json:"capture" binding:"required,max=10"`
}

type HarvestOperationDelete struct {
	UserId int
	Id     int `json:"id" binding:"required"`
}

type HarvestOperationsHist struct {
	Gradient []string `json:"gradient" binding:"required"`
	Hoids    []int    `json:"hoids" binding:"required"`
	Scale    string   `json:"scale" binding:"required"`
	Variable int      `json:"variable"` // binding:"required,gte=0"` //Problema con el 0
}

// Diferencia Pallete vs Gradient

type HarvestOperationsImg struct {
	Pallete  []string `form:"pallete" binding:"required"`
	Hoids    []int    `form:"hoids" binding:"required"`
	Scale    string   `form:"scale" binding:"required"`
	Variable int      `form:"variable"` // binding:"required,gte=0"` //Problema con el 0
}

//https://stackoverflow.com/questions/77934944/key-user-status-errorfield-validation-for-status-failed-on-the-required

type HarvestOperationsBounds struct {
	Hoids []int `json:"hoids" binding:"required"`
}

type HarvestOperationsStamps struct {
	X     float32
	Y     float32
	Value float32
	Color [3]int //[R,G,B]
}
