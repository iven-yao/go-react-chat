package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Key      string `json:"key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
