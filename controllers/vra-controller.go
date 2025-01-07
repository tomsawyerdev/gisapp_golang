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

/*
 VraList'
 VraCreate
 VraRename
 VraDelete
 VraChannelCreate
 VraChannelRename
 VraChannelUpdate
 VraChannelDelete
*/

//
//localhost:5000/vra/vralist
func VraList(c *gin.Context) {
	//userid := c.GetInt("userid") // return 0 if fail

	fmt.Println("VraList")
	//lastname := c.Query("long_name")

	//fmt.Println("RowsAffected:", result.RowsAffected, result2.RowsAffected)

	// get fieldid

	records, err := models.VraList(14)
	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "count": len(records), "items": records})
	//c.IndentedJSON(http.StatusOK, gin.H{"message": "GetPlayers"})//, "items": [1,2,3]
}

func VraCreate(c *gin.Context) {
	//fmt.Println("VraCreate")
	//userid := c.GetInt("userid")
	var requestBody dto.VraCreate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraCreate(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func VraRename(c *gin.Context) {
	fmt.Println("")
	//userid := c.GetInt("userid")
	var requestBody dto.VraRename

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraRename(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func VraDelete(c *gin.Context) {
	fmt.Println("")
	//userid := c.GetInt("userid")
	var requestBody dto.VraDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraDelete(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//----------------------------------
// Channels
//----------------------------------

func VraChannelCreate(c *gin.Context) {
	fmt.Println("VraChannelCreate")
	//userid := c.GetInt("userid")
	var requestBody dto.VraChannel

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraChannelCreate(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func VraChannelRename(c *gin.Context) {
	fmt.Println("VraChannelRename")
	//userid := c.GetInt("userid")
	var requestBody dto.VraChannelRename

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraChannelRename(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func VraChannelUpdate(c *gin.Context) {
	fmt.Println("VraChannelUpdate")
	//userid := c.GetInt("userid")
	var requestBody dto.VraChannelUpdate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraChannelUpdate(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func VraChannelDelete(c *gin.Context) {
	fmt.Println("VraChannelDelete")
	//userid := c.GetInt("userid")
	var requestBody dto.VraChannelUpdate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.VraChannelUpdate(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
