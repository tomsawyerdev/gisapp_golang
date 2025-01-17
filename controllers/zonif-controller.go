package controllers

import (
	"fmt"
	//svc "gisapi/database"
	"gisapi/dto"
	"gisapi/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	//"reflect"
	//"strconv"
)

/*
ZonifList
ZonifCreate
ZonifRename
ZonifUpdColors
ZonifDelete
ZonifCreateBuffer
ZoneRename
ZoneDelete
ZoneCreate
ZoneUpdBoundary
ZoneRemovePoints
ZoneSimplify
ZoneRefine
ZoneUpdClip
*/
func ZonifList(c *gin.Context) {

	fmt.Println("ZonifList")
	//lastname := c.Query("long_name")

	// Get fieldid

	var requestBody dto.ZonifList

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}

	records, err := models.ZonifList(requestBody.FieldId)
	//records, err := models.VraList(14)

	//fmt.Println("GetFarms, err:", err)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": 200, "count": len(records), "items": records})
}

func ZonifCreate(c *gin.Context) {
	var requestBody dto.ZonifCreate

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZonifCreate(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZonifRename(c *gin.Context) {
	//fmt.Println("")

	var requestBody dto.ZonifRename

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Fprintf(os.Stderr, "Failed BindJSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZonifRename(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})

}
func ZonifUpdColors(c *gin.Context) {
	var requestBody dto.ZonifUpdColors

	if err := c.BindJSON(&requestBody); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZonifUpdColors(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZonifDelete(c *gin.Context) {
	var requestBody dto.ZonifDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZonifDelete(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func ZonifCreateBuffer(c *gin.Context) {
	var requestBody dto.ZonifCreateBuffer

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Fprintf(os.Stderr, "Failed BindJSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZonifCreateBuffer(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// ---------------------------------
// Zone
// ---------------------------------

func ZoneRename(c *gin.Context) {
	var requestBody dto.ZoneRename

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Fprintf(os.Stderr, "Failed BindJSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZoneRename(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneDelete(c *gin.Context) {
	var requestBody dto.ZoneDelete

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZoneDelete(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// -----------------------
func ZoneCreate(c *gin.Context) {
	var requestBody dto.ZoneCreate

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Fprintf(os.Stderr, "Failed BindJSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	status, err := models.ZoneCreate(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	if status != nil {
		c.JSON(http.StatusOK, status)
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func ZoneUpdBoundary(c *gin.Context) {
	var requestBody dto.ZoneUpdBoundary

	if err := c.BindJSON(&requestBody); err != nil {
		fmt.Fprintf(os.Stderr, "Failed BindJSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	status, err := models.ZoneUpdBoundary(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	if status != nil {
		c.JSON(http.StatusOK, status)
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

//-----------------------

func ZoneRemovePoints(c *gin.Context) {
	var requestBody dto.ZoneRemovePoints

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZoneRemovePoints(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneSimplify(c *gin.Context) {
	var requestBody dto.ZoneSimplify

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZoneSimplify(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneRefine(c *gin.Context) {
	var requestBody dto.ZoneRefine

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZoneRefine(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneUpdClip(c *gin.Context) {
	var requestBody dto.ZoneUpdClip

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": 400, "message": "Invalid data"})
		return
	}
	//requestBody.Userid = c.GetInt("userid")
	//fmt.Println("requestBody:", requestBody)
	err := models.ZoneUpdClip(requestBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": 500, "error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
