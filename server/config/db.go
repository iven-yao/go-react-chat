package config

import (
	"chat-server/models"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PublicChatRoomID uint

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Chat{}, &models.Chatroom{}, &models.ChatroomMember{})
	DB = db

	Init()
}

func Init() {
	chatrooms := []models.Chatroom{}
	err := DB.Table("chatrooms").Find(&chatrooms).Error
	if err != nil {
		fmt.Println("select chatrooms panic")
		return
	}

	if len(chatrooms) == 0 {
		var defaultChatroom models.Chatroom
		defaultChatroom = models.Chatroom{
			Name: "Public",
		}
		DB.Table("chatrooms").Create(&defaultChatroom)
		PublicChatRoomID = defaultChatroom.ID
	}
}
