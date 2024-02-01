package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	Chatroom_id uint   `json:"chatroom_id"`
	User_id     uint   `json:"user_id"`
	Message     string `json:"message"`
	Upvotes     uint   `json:"upvotes"`
	Downvotes   uint   `json:"downvotes"`
}
