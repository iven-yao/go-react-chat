package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `json:"username" gorm:"unique; not null;type:varchar(100);default:null;"`
	Password string `json:"password" gorm:"not null;type:varchar(100);default:null;"`
}
