package main

//https://go.dev/doc/tutorial/web-service-gin#design_endpoints

import (
	"context"
	"fmt"

	//"strings"

	//"net/http"
	"os"

	"gisapi/controllers"
	"gisapi/database"

	//mismo nombre para la carpeta y el package, el archivo es: isAuth.go
	"gisapi/middlewares"
	"gisapi/routes"

	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// curl -v -X POST http://localhost:8080/example -H "content-type: application/json"
// curl http://localhost:8080/albums  -H "authorization:  Bearer  XCFRCVD"
// curl http://localhost:8080/api/albums  -H "authorization:  Bearer  XCFRCVD"

// go function optional paramete
// with midleware  handlers ...HandlerFunc

func AddRoutes(router *gin.Engine, path string, addRoutes func(*gin.RouterGroup)) {

	rgr := router.Group(path)
	//rgr.USE(middleware)
	//rgr.GET("/users", controllers.GetUsers)
	addRoutes(rgr)
}
func AddRoutes2(router *gin.Engine, path string, middleware gin.HandlerFunc, addRoutes func(*gin.RouterGroup)) {

	rgr := router.Group(path)
	rgr.Use(middleware)
	//rgr.GET("/users", controllers.GetUsers)
	addRoutes(rgr)

}

// CORS

func main() {

	exePath, err := os.Executable()
	if err != nil {
		fmt.Println("Failed to get executable path:", err)
	} else {
		fmt.Println("Executable path:", exePath)
	}
	// Init Database

	database.ConfigureDb()
	fmt.Println("Main Database Ping:", database.DB.Ping(context.Background()) == nil)
	//defer database.DB.Close(context.Background())
	defer database.DB.Close()

	// Init router

	router := gin.Default()

	// Init CORS  https://github.com/gin-contrib/cors

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://192.168.1.111:3000"}
	// Credential is not supported if the CORS header ‘Access-Control-Allow-Origin’ is ‘*’
	//config.AllowAllOrigins = true

	// (Reason: CORS Missing Allow Header)
	//config.AddAllowHeaders("*")
	config.AllowMethods = []string{"POST", "GET", "PUT", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization", "Accept", "User-Agent", "Cache-Control", "Pragma", "Session"}
	config.ExposeHeaders = []string{"Content-Length"}
	config.AllowCredentials = true
	config.MaxAge = 12 * time.Hour // 12*3600 = 43200

	fmt.Println("Hour:", 12*time.Hour, config.MaxAge)

	router.Use(cors.New(config))

	//router.Use(cors.Default())

	//router.Use(middlewares.Logger())

	AddRoutes(router, "/users", routes.SessionsRoutes)

	//AddRoutes2(router, "/farms", middlewares.SetId(), routes.FarmsRoutes)
	AddRoutes2(router, "/farms", middlewares.IsAuth(), routes.FarmsRoutes)
	AddRoutes2(router, "/fields", middlewares.IsAuth(), routes.FieldsRoutes)
	AddRoutes2(router, "/harvest", middlewares.IsAuth(), routes.HarvestRoutes)
	//AddRoutes2(router, "/zonif", middlewares.IsAuth(), routes.ZonifRoutes)
	AddRoutes2(router, "/zonif", middlewares.SetId(), routes.ZonifRoutes)

	AddRoutes2(router, "/vra", middlewares.IsAuth(), routes.VraRoutes)

	//AddRoutes(router, "/zonif", routes.ZonifRoutes)

	// Test routes
	router.GET("/json/:id", controllers.FieldGeoJson)

	//fmt.Printf("APi type: %T\n", rgr)

	//router.Run("localhost:5000")
	router.Run("0.0.0.0:5000")
}
