package models

import "gorm.io/gorm"

type Chatroom struct {
	gorm.Model
	Name string `json:"name"`
}
