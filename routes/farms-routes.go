package routes

import (
	"gisapi/controllers"
	"github.com/gin-gonic/gin"
)

//localhost:3000/farms/farms

func FarmsRoutes(rgr *gin.RouterGroup) {

	rgr.GET("/farms", controllers.FarmsList)
	//rgr.GET("/farmstree" do
	rgr.POST("/farmcreate", controllers.FarmCreate)
	rgr.POST("/farmupdate", controllers.FarmUpdate)
	rgr.POST("/farmdelete", controllers.FarmDelete)

}
