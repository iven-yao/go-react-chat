package controller

import (
	"chat-server/config"
	"chat-server/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	users := []models.User{}
	err := config.DB.Table("users").Find(&users).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	if len(users) == 0 {
		c.IndentedJSON(200, gin.H{"message": "empty"})
		return
	}
	c.IndentedJSON(200, &users)
}

func getAndHandleUserExists(user *models.User, username string) (exists bool, err error) {
	userExistsQuery := config.DB.Table("users").Where("username = ?", strings.ToLower(username)).Limit(1).Find(&user)

	if userExistsQuery.Error != nil {
		return false, userExistsQuery.Error
	}

	userExists := userExistsQuery.RowsAffected > 0

	if userExists == true {
		return true, nil
	}

	return false, nil
}

func hashAndSaltPassword(password string) (hashedPassword string, err error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}

func Register(c *gin.Context) {
	var userInput models.User
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var newUser models.User
	userExists, userExistsErr := getAndHandleUserExists(&newUser, userInput.Username)
	if userExistsErr != nil {
		c.JSON(http.StatusBadRequest, userExistsErr.Error())
		return
	}

	if userExists == true {
		c.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	}

	hashedPassword, hashedErr := hashAndSaltPassword(userInput.Password)
	if hashedErr != nil {
		c.JSON(http.StatusBadRequest, hashedErr.Error())
		return
	}

	newUser = models.User{
		Username: strings.ToLower(userInput.Username),
		Password: hashedPassword,
	}

	config.DB.Table("users").Create(&newUser)
	c.IndentedJSON(200, gin.H{
		"ID":       newUser.ID,
		"Username": newUser.Username,
	})
}

func Login(c *gin.Context) {
	var userInput models.User
	err := c.BindJSON(&userInput)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	var existingUser models.User
	errorMsg := "Invalid username or password"
	userExists, userExistsErr := getAndHandleUserExists(&existingUser, userInput.Username)
	if userExistsErr != nil {
		c.JSON(http.StatusInternalServerError, userExistsErr.Error())
		return
	}

	if userExists == false {
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMsg})
		return
	}

	passwordErr := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(userInput.Password))
	if passwordErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": errorMsg})
		return
	}

}
