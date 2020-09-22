package main

import (
	"go-server/router"
)

func main() {
	// database.InitMysql()
	// defer database.Db.Close()
	r := router.Create()
	r.Run(":8080")
}
