package routes

import (
	"gisapi/controllers"
	"github.com/gin-gonic/gin"
)

//localhost:3000/farms/farms

func FieldsRoutes(rgr *gin.RouterGroup) {

	//rgr.GET("/fieldtest", controllers.TestField)

	rgr.POST("/fields", controllers.FieldsList)       //FieldsList
	rgr.POST("/fieldcreate", controllers.FieldCreate) //FieldsNew
	rgr.POST("/fieldrename", controllers.FieldUpdate)
	rgr.POST("/fielddelete", controllers.FieldDelete)
	rgr.POST("/fieldboundary", controllers.FieldBoundary)
	rgr.GET("/fielddownloadgeojson/:id", controllers.FieldGeoJson)

}
