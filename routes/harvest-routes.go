package routes

import (
	"gisapi/controllers"
	"github.com/gin-gonic/gin"
)

//localhost:3000/farms/farms

func HarvestRoutes(rgr *gin.RouterGroup) {

	rgr.POST("/harvestlist", controllers.HarvestList)

	// -------------------------------
	// Seasons
	//--------------------------------

	rgr.POST("/harvestseasoncreate", controllers.HarvestSeasonCreate)
	rgr.POST("/harvestseasonupdate", controllers.HarvestSeasonUpdate)
	rgr.POST("/harvestseasondelete", controllers.HarvestSeasonDelete)

	// -------------------------------
	// Operations
	//--------------------------------

	rgr.POST("/harvestoperationcreate", controllers.HarvestOperationCreate)
	rgr.POST("/harvestoperationupdate", controllers.HarvestOperationUpdate)
	rgr.POST("/harvestoperationdelete", controllers.HarvestOperationDelete)
	//rgr.POST("/harvestimportlogcsv"

	// ----------------------------------------------------------------------
	// Operations Histogram
	// localhost:5000/harvest/harvestoperationshist
	// ----------------------------------------------------------------------

	rgr.POST("/harvestoperationshist", controllers.HarvestOperationsHist)
	rgr.GET("/harvestoperationshist", controllers.HarvestOperationsHist) // Only for test

	rgr.POST("/harvestoperationbounds", controllers.HarvestOperationBounds)
	rgr.GET("/harvestoperationbounds", controllers.HarvestOperationBounds) // Only for test

	//rgr.POST("/harvestoperationimg", controllers.HarvestOperationImg)
	rgr.GET("/harvestoperationimg", controllers.HarvestOperationImg)

	//rgr.GET("/test_receive_array", controllers.receive_array)// Only for test
	rgr.GET("/test_colors", controllers.TestColors)       // Only for test
	rgr.GET("/test_image", controllers.TestImageCreation) // Only for test

}
