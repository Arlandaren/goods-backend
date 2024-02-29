package handlers

import (
	"net/http"
	"server/pkg/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListGoods(c *gin.Context){
	limit,_ := strconv.Atoi(c.Query("limit"))
	offset,_ := strconv.Atoi(c.Query("offset"))
	if limit == 0{
		limit = 10
	}
	if offset == 0{
		offset = 1
	}
	goods,meta,err := models.GetGoodsFromDB(limit,offset)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error with db", "details":err.Error()})
		return
	}
	c.JSON(http.StatusOK,gin.H{"meta":meta,"goods":goods})
}