package config

import (
	"chat-server/models"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func goDotEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println((err.Error()))
		return "localhost"
	}

	return os.Getenv(key)
}

func Connect() {
	db_host := goDotEnv("DB_HOST")
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@"+db_host+":5432/postgres"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{}, &models.Chat{})
	DB = db
}
