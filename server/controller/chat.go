package controller

import (
	"chat-server/config"
	"chat-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type result struct {
	gorm.Model
	Username  string `json:"username"`
	Message   string `json:"message"`
	Upvotes   uint   `json:"upvotes"`
	Downvotes uint   `json:"downvotes"`
}

func GetChats(c *gin.Context) {
	userid := getUserId(c)
	print(userid)

	chatroomId := c.Param("chatroomId")

	chats := []result{}

	err := config.DB.Table("chats").Select(
		"chats.id, users.username, chats.message, chats.upvotes, chats.downvotes, chats.created_at").Joins(
		"JOIN users ON chats.user_id = users.id").Where(
		"chatroom_id = ?", chatroomId).Order(
		"created_at ASC").Find(&chats).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(200, &chats)
}

func SendChat(c *gin.Context) {
	userid := getUserId(c)

	var userInput models.Chat
	if err := c.BindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	// check input
	if userInput.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "empty input"})
		return
	}

	var newChat models.Chat
	newChat = models.Chat{
		Chatroom_id: 1,
		User_id:     uint(userid),
		Message:     userInput.Message,
		Downvotes:   0,
		Upvotes:     0,
	}

	config.DB.Table("chats").Create(&newChat)
	c.IndentedJSON(200, gin.H{"message": "message recved"})
}
