package routes

import (
	"gisapi/controllers"
	"github.com/gin-gonic/gin"
)

// cuidado con session/ vs session

func SessionsRoutes(rgr *gin.RouterGroup) {

	rgr.GET("", controllers.SessionGet)
	rgr.POST("/login", controllers.SessionCreate) //controllers.NewSession; controllers.sessions.new

}
