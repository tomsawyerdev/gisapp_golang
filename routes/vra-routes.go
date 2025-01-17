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
	rgr.POST("/vracreate", controllers.VraCreate)
	rgr.POST("/vrarename", controllers.VraRename)
	rgr.POST("/vradelete", controllers.VraDelete)
	//Channels:
	rgr.POST("/vrachannelcreate", controllers.VraChannelCreate)
	rgr.POST("/vrachannelrename", controllers.VraChannelRename)
	rgr.POST("/vrachannelupdate", controllers.VraChannelUpdate)
	rgr.POST("/vrachanneldelete", controllers.VraChannelDelete)

}
