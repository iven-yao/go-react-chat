package controller

import (
	"chat-server/config"
	"chat-server/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	config.DB.Find(&users)
	c.IndentedJSON(200, &users)
}

func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.BindJSON(&user); err != nil {
		return
	}

	config.DB.Create(&user)
	c.IndentedJSON(200, &user)
}

func DeleteUser(c *gin.Context) {
	var user models.User
	var id string = c.Param("id")
	config.DB.Where("id = ?", id).Delete(&user)
	c.IndentedJSON(200, gin.H{"message": "user" + id + " is deleted"})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.BindJSON(&user)
	config.DB.Save(&user)
	c.IndentedJSON(200, &user)
}

func GetUserById(c *gin.Context) {
	var user models.User
	config.DB.Where("id = ?", c.Param("id")).First(&user)
	c.IndentedJSON(200, &user)
}
