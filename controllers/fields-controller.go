package controllers

import (
	"fmt"
	//svc "gisapi/database"
	"gisapi/dto"
	"gisapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	//"reflect"
	"strconv"
)

// http.StatusBadRequest 400
//
//localhost:3000/farms/farms

// func FieldsList(c *gin.Context) {
func FieldsList(c *gin.Context) {
	userid := c.GetInt("userid") // return 0 if fail

	fmt.Println("ListFields ------------")
	fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.FieldList

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	records, err := models.FieldsList(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "items": records})
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "GetPlayers"})//, "items": [1,2,3]
}

func FieldCreate(c *gin.Context) {

	userid := c.GetInt("userid") // return 0 if fail

	fmt.Println("NewField ------------")
	fmt.Println("   userid:", userid)

	var requestBody dto.FieldCreate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	// TODO: Verify ownership

	err := models.FieldVerifyOwnership(requestBody)
	//fmt.Println("   FieldVerifyOwnership:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
		return
	}

	//------------------------

	//fmt.Println("   requestBody:", requestBody)

	// Create

	if requestBody.Type == "circle" {

		if requestBody.Radius < 100 {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid Radius"})
			return
		}

		err := models.FieldcreateCircle(requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return

	}

	//-------------------------
	// Create a Polygon
	//-------------------------
	// Check for Validity
	//-------------------------
	//fmt.Println("   requestBody.Polygon:", requestBody.Polygon)

	verify, reason, err := models.FieldcreateVerifyBoundary(requestBody.Polygon)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
		return
	}

	if !verify {
		c.JSON(http.StatusOK, gin.H{"status": "error", "msg": reason})
	}

	//-------------------------
	// Create Polygon
	//-------------------------
	//
	err = models.FieldcreatePolygon(requestBody)
	if err != nil {
		fmt.Println(" FieldcreatePolygon:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

// Rename Field
func FieldUpdate(c *gin.Context) {
	userid := c.GetInt("userid") // return 0 if fail

	fmt.Println("Rename Field ------------")
	fmt.Println("   userid:", userid)

	var requestBody dto.FieldRename

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	// TODO: Verify ownership

	err := models.Fieldrename(requestBody)
	//fmt.Println("   FieldVerifyOwnership:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func FieldDelete(c *gin.Context) {

	var requestBody dto.FieldDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	// TODO: Verify ownership

	err := models.Fielddelete(requestBody)
	//fmt.Println("   FieldVerifyOwnership:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Update Boundary
func FieldBoundary(c *gin.Context) {

	var requestBody dto.FieldBoundary

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Println("   BndField err:", err)

		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	// TODO: Verify ownership
	/*
		err := models.FieldVerifyOwnership(requestBody)
		//fmt.Println("   FieldVerifyOwnership:", err)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}*/

	//------------------------

	//fmt.Println("   requestBody:", requestBody)

	// Create

	if requestBody.Type == "circle" {

		if requestBody.Radius < 100 {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid Radius"})
			return
		}

		err := models.FieldboundaryCircle(requestBody)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success"})
		return

	}

	//-------------------------
	// Update a Polygon
	//-------------------------
	// Check for Validity
	//-------------------------
	//fmt.Println("   requestBody.Polygon:", requestBody.Polygon)

	verify, reason, err := models.FieldcreateVerifyBoundary(requestBody.Polygon)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
		return
	}

	if !verify {
		c.JSON(http.StatusOK, gin.H{"status": "error", "msg": reason})
	}

	//-------------------------
	// Create Polygon
	//-------------------------
	//
	err = models.FieldboundaryPolygon(requestBody)
	if err != nil {
		fmt.Println(" FieldboundaryPolygon:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}

// Export Geojson
func FieldGeoJson(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid ID"})
		return
	}

	//fmt.Println("   GeoField:", id)

	geojson, _ := models.Fielddownloadgeojson(id)

	//fmt.Println("   GeoField:", geojson)

	c.Header("Content-Disposition", "attachment; filename=field.json")
	//c.Header("Content-Type", "text/json")
	c.String(http.StatusOK, geojson)
}

//localhost:5000/fields/fieldtest

func FieldTest(c *gin.Context) {

	var polygonMap = make(map[string]any)
	polygonMap["type"] = "Polygon"

	coordinates := [][][]float32{{{-63.807564, -31.303782}, {-63.774948, -31.305542}, {-63.768768, -31.332525}, {-63.801041, -31.336337}, {-63.807564, -31.303782}}}
	polygonMap["coordinates"] = coordinates

	//polygon := `{"type":"Polygon","coordinates":[[[-63.846016,-31.299968],[-63.830223,-31.296155],[-63.800011,-31.328419],[-63.83503,-31.339856],[-63.846016,-31.299968]]]}`
	//Badpolygon := `{"type":"Polygon","coordinates":[[[-63.846016,-31.299968],[-63.800011,-31.328419],[-63.830223,-31.296155],[-63.83503,-31.339856],[-63.846016,-31.299968]]]}`

	verify, reason, err := models.FieldcreateVerifyBoundary(polygonMap) //requestBody.Polygon)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "message": err})
		return
	}

	fmt.Println("   TestField:", verify, reason)
	c.JSON(http.StatusOK, gin.H{"status": "success", "verify": verify, "reason": reason})

}
