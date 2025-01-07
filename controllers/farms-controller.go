package controllers

import (
	"fmt"
	//svc "gisapi/database"
	"gisapi/dto"
	"gisapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	//"reflect"
	//"strconv"
)

// http.StatusBadRequest 400
//
//localhost:3000/farms/farms
func FarmsList(c *gin.Context) {
	userid := c.GetInt("userid") // return 0 if fail

	fmt.Println("GetFarms, userid:", userid)
	//lastname := c.Query("long_name")

	//fmt.Println("RowsAffected:", result.RowsAffected, result2.RowsAffected)
	records, err := models.FarmsList(userid)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "count": len(records), "items": records})
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "GetPlayers"})//, "items": [1,2,3]
}

// farmcreate

func FarmCreate(c *gin.Context) {
	fmt.Println("NewFarm")
	//userid := c.GetInt("userid")
	var requestBody dto.FarmNew

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	requestBody.Userid = c.GetInt("userid")

	//fmt.Println("requestBody:", requestBody)

	err := models.FarmCreate(requestBody) //Que le paso
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// Arreglar las validaciones en el modelo `binding:"required"`
func FarmUpdate(c *gin.Context) {

	//fmt.Println("UpdFarm")
	//userid := c.GetInt("userid")
	var requestBody dto.FarmNew

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	requestBody.Userid = c.GetInt("userid")

	//fmt.Println("requestBody:", requestBody)

	err := models.FarmUpdate(requestBody) //Que le paso
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func FarmDelete(c *gin.Context) {

	var requestBody dto.FarmNew

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	requestBody.Userid = c.GetInt("userid")

	fmt.Println("Del farm, requestBody:", requestBody)

	err := models.FarmDelete(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
