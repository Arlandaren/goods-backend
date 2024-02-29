package main

import (
	"server/pkg/router"

	"github.com/gin-gonic/gin"
)

func main(){
	r := gin.Default()
	router.RouteAll(r)
	r.Run()
}