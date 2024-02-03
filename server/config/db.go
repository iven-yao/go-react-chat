package config

import (
	"chat-server/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB
var PublicChatRoomID uint

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@db:5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Chat{})
	DB = db
}
