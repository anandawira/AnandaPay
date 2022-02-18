package model

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
}
