package models

import "gorm.io/gorm"

type Chat struct {
	gorm.Model
	User_id   uint   `json:"user_id"`
	Message   string `json:"message"`
	Upvotes   uint   `json:"upvotes"`
	Downvotes uint   `json:"downvotes"`
}
