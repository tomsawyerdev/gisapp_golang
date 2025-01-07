package routes

import (
	"gisapi/controllers"
	"github.com/gin-gonic/gin"
)

//localhost:3000/farms/farms

//----------------------------------------------------------------------
// Variable Rate Applications
//----------------------------------------------------------------------
/*
 /vralist'
 /vracreate
 /vrarename
 /vradelete
 /vrachannelcreate
 /vrachannelrename
 /vrachannelupdate
 /vrachanneldelete
*/

func VraRoutes(rgr *gin.RouterGroup) {

	rgr.POST("/vralist", controllers.VraList)
	rgr.GET("/vralist", controllers.VraList) //for test
	/*
		rgr.POST("/vracreate", controllers.Vracreate)
		rgr.POST("/vrarename", controllers.Vrarename)
		rgr.POST("/vradelete", controllers.Vradelete)
		rgr.POST("/vrachannelcreate", controllers.Vrachannelcreate)
		rgr.POST("/vrachannelrename", controllers.Vrachannelrename)
		rgr.POST("/vrachannelupdate", controllers.Vrachannelupdate)
		rgr.POST("/vrachanneldelete", controllers.Vrachanneldelete)
	*/

}
