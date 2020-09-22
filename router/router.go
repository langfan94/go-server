package router

import (
	"go-server/api"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func Create() *gin.Engine {
	// Logging to a file.
	f, _ := os.Create("gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/", api.All)
		v1.GET("/login", api.FindAll)
		v1.GET("/user/:name/:id", api.User)
	}

	return router
}
