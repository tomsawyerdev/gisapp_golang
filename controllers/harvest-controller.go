package controllers

import (
	"bytes"
	"fmt"
	//svc "gisapi/database"
	//"encoding/json"
	"gisapi/colors"
	"gisapi/dto"
	"gisapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	//"strco/nv"
	"image/png"
	"time"
)

// http.StatusBadRequest 400
//
//localhost:3000/farms/farms

func HarvestList(c *gin.Context) {
	//fmt.Println("HarvestList ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestList

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	records, err := models.HarvestList(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "items": records})
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "GetPlayers"})//, "items": [1,2,3]
}

/*
HarvestSeasonCreate
HarvestSeasonUpdate
HarvestSeasonDelete
HarvestOperationCreate
HarvestOperationUpdate
HarvestOperationDelete
HarvestOperationsHist
HarvestOperationBounds
HarvestOperationImg*/

func HarvestSeasonCreate(c *gin.Context) {
	//fmt.Println("HarvestSeasonCreate ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestSeasonCreate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	err := models.HarvestSeasonCreate(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func HarvestSeasonUpdate(c *gin.Context) {

	//fmt.Println("HarvestSeasonUpdate ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestSeasonUpdate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	err := models.HarvestSeasonUpdate(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func HarvestSeasonDelete(c *gin.Context) {

	//fmt.Println("HarvestSeasonDelete ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestSeasonDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	err := models.HarvestSeasonDelete(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func HarvestOperationCreate(c *gin.Context) {

	//fmt.Println("HarvestOperationCreate ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestOperationCreate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	err := models.HarvestOperationCreate(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func HarvestOperationUpdate(c *gin.Context) {

	//fmt.Println("HarvestOperationUpdate ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestOperationUpdate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	err := models.HarvestOperationUpdate(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func HarvestOperationDelete(c *gin.Context) {

	//fmt.Println("HarvestOperationDelete ------------")
	//userid := c.GetInt("userid") // return 0 if fail
	//fmt.Println("   userid:", userid)
	//lastname := c.Query("long_name")
	var requestBody dto.HarvestOperationDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.UserId = c.GetInt("userid")

	//fmt.Println("   requestBody:", requestBody)
	err := models.HarvestOperationDelete(requestBody)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// localhost:5000/harvest/harvestoperationshist
func HarvestOperationsHist(c *gin.Context) {
	fmt.Println("HarvestOperationsHist ------------")

	var requestBody dto.HarvestOperationsHist
	//fmt.Println("c.Request.Method", c.Request.Method)

	if c.Request.Method == "GET" {
		requestBody = dto.HarvestOperationsHist{Hoids: []int{1}, Variable: 0, Scale: "I4", Gradient: []string{"#FF0000", "#FFD700", "#006455"}}
	} else {

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			//fmt.Println("Error:", err)
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
			return
		}

	}
	//fmt.Println("requestBody:", requestBody)
	//fmt.Println("requestBody:", requestBody.Variable)

	values, err := models.HarvestOperationsValues(requestBody)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	// Generate Histogram Values,

	result := models.HarvestGenerateHistogram(requestBody.Gradient, requestBody.Scale, values)

	c.JSON(http.StatusOK, result)
	//c.JSON(http.StatusOK, gin.H{"status": 200, "values": values})
}

// Return Bound in 4326
// {"latmax":-31.391349786,"latmin":-31.394299467,"lonmax":-63.690877312,"lonmin":-63.69499961}
// localhost:5000/harvest/harvestoperationbounds  --> e singular
func HarvestOperationBounds(c *gin.Context) {
	fmt.Println("HarvestOperationBounds ------------4326")

	var requestBody dto.HarvestOperationsBounds
	//fmt.Println("c.Request.Method", c.Request.Method)

	if c.Request.Method == "GET" {
		requestBody = dto.HarvestOperationsBounds{Hoids: []int{3, 4}}
	} else {

		if err := c.BindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
			return
		}

	}
	//fmt.Println("requestBody:", requestBody)

	values, err := models.HarvestOperationsBounds4326(requestBody.Hoids)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "bounds": values})
}

// http://localhost:5000/harvest/harvestoperationimg?hoids=1&variable=2&scale=I3
// http://localhost:5000/harvest/harvestoperationimg?noise=1734364954991&scale=I3&hoids=1&variable=0&pallete=%23FF0000&pallete=%23FFD700&pallete=%23006400"
//
//	http://localhost:5000/harvest/harvestoperationimg?noise=1734470271719&scale=I3&hoids=1&variable=0&pallete=%23FF0000&pallete=%23FFD700&pallete=%23006400%22
func HarvestOperationImg(c *gin.Context) {

	fmt.Println("HarvestOperationIMG ------------")

	var requestBody dto.HarvestOperationsImg
	//fmt.Println("c.Request.Method", c.Request.Method)

	//Siempre es GET

	if c.Request.Method == "POST" {
		c.JSON(http.StatusOK, gin.H{"status": 400, "message": "Post not allowed"})
	}

	if err := c.BindQuery(&requestBody); err != nil {

		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		//fmt.Println("Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data", "error": err})
		return

	}

	fmt.Println("RequestBody:", requestBody)

	//requestBody = dto.HarvestOperationsImg{Hoids: []int{3, 4}, Variable: 0, Scale: "I4", Gradient: []string{"#FF0000", "#FFD700", "#006455"}}
	// Enlapses time --------------
	start := time.Now()
	fmt.Println(start)

	img := models.HarvestOperationsImage(requestBody)

	//return send_file(imgbuffer, mimetype='image/jpeg')

	buffer := new(bytes.Buffer)
	png.Encode(buffer, img)
	c.Writer.Write(buffer.Bytes())

}

// localhost:5000/harvest/test_colors
func TestColors(c *gin.Context) {
	fmt.Println("Harvest  TestColors ------------")

	colors.TestColors()

}

// localhost:5000/harvest/test_image
func TestImageCreation(c *gin.Context) {
	fmt.Println("Harvest  TestImageCreation ------------")

	new_image := models.TestImageCreation()
	buffer := new(bytes.Buffer)
	err := png.Encode(buffer, new_image)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	//c.Header("Content-Type", "image/png")
	//c.Header("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	//c.Writer.WriteHeader(http.StatusOK)
	c.Writer.Write(buffer.Bytes())

}
