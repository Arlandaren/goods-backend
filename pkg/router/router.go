package router

import (
	"server/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RouteAll(r *gin.Engine){
	api := r.Group("api")
	{
		good := api.Group("goods")
		{
			good.GET("list",handlers.ListGoods)
		}
	}
}