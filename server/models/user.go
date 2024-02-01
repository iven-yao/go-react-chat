package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       string `json:"ID" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
