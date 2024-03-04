package handlers

import (
	"encoding/json"
	"net/http"
	"server/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListGoods(c *gin.Context){
	limit,err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	offset,err := strconv.Atoi(c.DefaultQuery("offset","1"))
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	if limit == 0{
		limit = 10
	}
	if offset == 0{
		offset = 1
	}
	data,err := models.GetGoods(limit,offset)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error with db", "details":err.Error()})
		return
	}
	var goods map[string]interface{}

    if err := json.Unmarshal(data, &goods); err != nil {
        c.JSON(http.StatusInternalServerError,gin.H{"error":"error with json","details":err.Error()})
		return
    }

	c.JSON(http.StatusOK,goods)
}
func RemoveGood(c *gin.Context){
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	projectId, err := strconv.Atoi(c.Query("projectId"))
	if err != nil {
		c.JSON(400, gin.H{"reason": err.Error()})
		return
	}
	
	if err := models.RemoveGood(id,projectId); err !=nil{
		c.JSON(404, gin.H{"reason":err.Error()})
		return
	}
	c.JSON(200,gin.H{"id":id,"projectId":projectId,"removed":true})
}
func CreateGood(c *gin.Context){
	projectId, err := strconv.Atoi(c.Query("projectId"))
	if err != nil {
		c.JSON(400, gin.H{"reason": "несоответствие формату"})
		return
	}
	var goodrequest *models.CreateRequest
	if err:=c.ShouldBindJSON(&goodrequest); err !=nil{
		c.JSON(400, gin.H{"reason": "несоответствие формату"})
		return
	}
	good,err := models.CreateGoodToDB(goodrequest.Name,projectId)
	if err != nil{
		c.JSON(404, gin.H{"reason": err.Error()})
		return
	}
	 c.JSON(200,good)
}