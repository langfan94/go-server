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

type Person struct {
	ID   string `uri:"id" binding:"required,uuid"`
	Name string `uri:"name" binding:"required"`
}

func User(c *gin.Context) {
	var person Person
	if err := c.ShouldBindUri(&person); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	c.JSON(200, gin.H{"name": person.Name, "uuid": person.ID})
}
