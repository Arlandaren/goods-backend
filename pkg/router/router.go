package router

import "github.com/gin-gonic/gin"

func RouteAll(r *gin.Engine){
	api := r.Group("api")
	{
		api.GET("")
	}
}