package api

import (
	"fmt"
	"go-server/database"
	"go-server/models"

	"github.com/gin-gonic/gin"
)

func FindAll(c *gin.Context) {
	rows, err := database.Db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error())
	}

	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Phone, &user.Description)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(user)
		users = append(users, user)
	}
	c.JSON(200, gin.H{
		"status": 200, "message": users,
	})
}

func All(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
