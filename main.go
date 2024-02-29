package main

import (
	"os"
	"server/pkg/models"
	"server/pkg/router"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main(){
	Logger := NewLogger()
	gin.DefaultWriter = Logger.Writer()
	err := godotenv.Load()
	if err != nil{
		Logger.Fatal("Error loading env variables")
	}
	err = models.InitDB(os.Getenv("POSTGRES_CONN"))
	if err != nil{
		Logger.Fatal(err.Error())
	}
	// models.InitRedis(os.Getenv("REDIS_CONN"))
	r := gin.Default()
	router.RouteAll(r)
	r.Run(os.Getenv("SERVER_ADDRESS"))
}