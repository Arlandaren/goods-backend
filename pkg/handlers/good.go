package handlers

import (
	"encoding/json"
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
	data,err := models.GetGoods(limit,offset)
	if err !=nil{
		c.JSON(http.StatusInternalServerError,gin.H{"error":"error with db", "details":err.Error()})
		return
	}
	var goods map[string]interface{}

    if err := json.Unmarshal(data, &goods); err != nil {
        c.JSON(http.StatusInternalServerError,gin.H{"error":"error with json","details":err.Error()})
    }

	c.JSON(http.StatusOK,goods)
}