package models

import "gorm.io/gorm"

type ChatroomMember struct {
	gorm.Model
	Chatroom_id uint `json:"chatroom_id"`
	User_id     uint `json:"user_id"`
}
