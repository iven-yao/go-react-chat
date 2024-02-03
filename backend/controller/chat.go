package controller

import (
	"chat-server/config"
	"chat-server/models"
	"encoding/json"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type result struct {
	gorm.Model
	ID        uint      `json:"ID"`
	CreatedAt time.Time `json:"CreatedAt"`
	Username  string    `json:"username"`
	Message   string    `json:"message"`
	Upvotes   uint      `json:"upvotes"`
	Downvotes uint      `json:"downvotes"`
	Type      string    `json:"type"`
}

type wsMessage struct {
	Type    string `json:"type"`
	Message string `json:"message"`
	Token   string `json:"token"`
	Vote    int    `json:"vote"`
	Chatid  uint   `json:"chatid"`
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var clients []websocket.Conn
var m sync.Mutex

func handleNewMessage(wsMsg wsMessage, userid uint, msgType int) {
	// check input
	if wsMsg.Message == "" {
		fmt.Println("empty input, abort")
	}

	// get username
	var username string
	config.DB.Table("users").Select("users.username").Where("users.id = ?", userid).Find(&username)

	newChat := models.Chat{
		User_id:   userid,
		Message:   wsMsg.Message,
		Downvotes: 0,
		Upvotes:   0,
	}

	config.DB.Table("chats").Create(&newChat)
	broadcastMsg := result{Username: username, Message: wsMsg.Message, Upvotes: newChat.Upvotes, Downvotes: newChat.Downvotes, ID: newChat.ID, CreatedAt: newChat.CreatedAt, Type: "NEWMESSAGE"}
	jsonBroadcast, err := json.Marshal(broadcastMsg)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, client := range clients {
		if err = client.WriteMessage(msgType, jsonBroadcast); err != nil {
		}
	}
}

func handleNewVote(wsMsg wsMessage, msgType int) {
	// get existing chat
	var existChat models.Chat
	if err := config.DB.Table("chats").Where("id = ?", wsMsg.Chatid).Find(&existChat).Error; err != nil {
		fmt.Println(err.Error())
		return
	}
	// get username
	var username string
	config.DB.Table("users").Select("users.username").Where("users.id = ?", existChat.User_id).Find(&username)

	if wsMsg.Vote > 0 {
		// upvote
		existChat.Upvotes = existChat.Upvotes + 1
		fmt.Println("upvote:", existChat.Upvotes)
		if err := config.DB.Table("chats").Where("id = ?", wsMsg.Chatid).Update("upvotes", existChat.Upvotes).Error; err != nil {
			fmt.Println(err.Error())
		}
	} else if wsMsg.Vote < 0 {
		// downvote
		existChat.Downvotes = existChat.Downvotes + 1
		fmt.Println("downvote:", existChat.Downvotes)
		if err := config.DB.Table("chats").Where("id = ?", wsMsg.Chatid).Update("downvotes", existChat.Downvotes).Error; err != nil {
			fmt.Println(err.Error())
		}
	}

	broadcastMsg := result{Username: username, Message: existChat.Message, Upvotes: existChat.Upvotes, Downvotes: existChat.Downvotes, ID: existChat.ID, CreatedAt: existChat.CreatedAt, Type: "VOTE"}
	jsonBroadcast, err := json.Marshal(broadcastMsg)
	if err != nil {
		fmt.Println(err.Error())
	}

	for _, client := range clients {
		if err = client.WriteMessage(msgType, jsonBroadcast); err != nil {
		}
	}

}

func WebSocketConnect(c *gin.Context) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)

	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	clients = append(clients, *ws)

	for {
		msgType, msg, err := ws.ReadMessage()
		if err != nil {
			return
		}

		fmt.Printf("%s sent: %s\n", ws.RemoteAddr(), string(msg))

		// decode the message
		var wsMsg wsMessage
		jsonErr := json.Unmarshal(msg, &wsMsg)
		if jsonErr != nil {
			fmt.Println(jsonErr.Error())
			return
		}

		// verify token
		userid, verifyErr := verifyToken(wsMsg.Token)
		if verifyErr != nil {
			fmt.Println(verifyErr.Error())
			return
		}

		if wsMsg.Type == "MESSAGE" {
			handleNewMessage(wsMsg, uint(userid), msgType)
		} else if wsMsg.Type == "VOTE" {
			handleNewVote(wsMsg, msgType)
		}
	}
}

func GetChats(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")[len("Bearer "):]
	_, verifyErr := verifyToken(token)
	if verifyErr != nil {
		fmt.Print("Invalid token")
		c.JSON(http.StatusBadRequest, verifyErr.Error())
		return
	}

	chats := []result{}

	err := config.DB.Table("chats").Select(
		"chats.id, users.username, chats.message, chats.upvotes, chats.downvotes, chats.created_at").Joins(
		"JOIN users ON chats.user_id = users.id").Order(
		"created_at ASC").Find(&chats).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.IndentedJSON(200, &chats)
}
