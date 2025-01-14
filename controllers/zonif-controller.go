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
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZonifRename(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZonifUpdColors(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZonifDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func ZonifCreateBuffer(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

// ---------------------------------
// Zone
// ---------------------------------

func ZoneRename(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneDelete(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneCreate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneUpdBoundary(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneRemovePoints(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneSimplify(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneRefine(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
func ZoneUpdClip(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "success"})
}
