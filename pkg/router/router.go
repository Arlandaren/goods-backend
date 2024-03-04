package router

import (
	"server/pkg/handlers"

	"github.com/gin-gonic/gin"
)

func RouteAll(r *gin.Engine){
	api := r.Group("api")
	{
		goods := api.Group("goods")
		{
			goods.GET("list",handlers.ListGoods)
		}
		good := api.Group("good")
		{
			good.DELETE("/remove", handlers.RemoveGood)
			good.POST("create", handlers.CreateGood)
		}
	}

}