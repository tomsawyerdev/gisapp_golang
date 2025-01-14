package routes

import (
	"gisapi/controllers"
	"github.com/gin-gonic/gin"
)

//localhost:3000/farms/farms

//----------------------------------------------------------------------
// Zonifications
//----------------------------------------------------------------------

func ZonifRoutes(rgr *gin.RouterGroup) {

	rgr.POST("/zoniflist", controllers.ZonifList)

	rgr.POST("/zonifcreate", controllers.ZonifCreate)
	rgr.POST("/zonifrename", controllers.ZonifRename)
	rgr.POST("/zonifupdcolors", controllers.ZonifUpdColors)
	rgr.POST("/zonifdelete", controllers.ZonifDelete)
	rgr.POST("/zonifcreatebuffer", controllers.ZonifCreateBuffer)
	rgr.POST("/zonerename", controllers.ZoneRename)
	rgr.POST("/zonedelete", controllers.ZoneDelete)
	rgr.POST("/zonecreate", controllers.ZoneCreate)
	rgr.POST("/zoneupdboundary", controllers.ZoneUpdBoundary)
	rgr.POST("/zoneremovepoints", controllers.ZoneRemovePoints)
	rgr.POST("/zonesimplify", controllers.ZoneSimplify)
	rgr.POST("/zonerefine", controllers.ZoneRefine)
	rgr.POST("/zoneupdclip", controllers.ZoneUpdClip)

}
