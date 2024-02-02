package controller

import (
	"chat-server/config"
	"chat-server/models"
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsers(c *gin.Context) {
	users := []User{}
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

func TestUser(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")[len("Bearer "):]
	id, err := verifyToken(token)
	if err != nil {
		fmt.Print("Invalid token")
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
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

func Register(c *gin.Context) {
	var userInput models.User
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check input
	if userInput.Password == "" || userInput.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "empty input"})
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

	// add to chatroom_members
	var defaultChatroomMember models.ChatroomMember
	defaultChatroomMember = models.ChatroomMember{
		Chatroom_id: config.PublicChatRoomID,
		User_id:     newUser.ID,
	}
	config.DB.Table("chatroom_members").Create(&defaultChatroomMember)

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

	token, tokenErr := createToken(existingUser.Username, existingUser.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, tokenErr.Error())
		return
	}

	c.IndentedJSON(200, gin.H{"message": "Login sucessful", "token": token, "ID": existingUser.ID, "username": existingUser.Username})

}
