package main

import (
	"go-server/api"
	"go-server/database"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitMysql()
	defer database.Db.Close()
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/", api.All)
		v1.GET("/login", api.FindAll)
	}

	router.Run(":8080")
}
