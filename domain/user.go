package domain

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `json:"fullname" gorm:"not null"`
	Email          string `json:"email" gorm:"unique;not null"`
	HashedPassword string `json:"password" gorm:"not null"`
	IsVerified     bool   `json:"isVerified" gorm:"not null"`
}

type UserUsecase interface {
	Register(fullname, email, plainPassword string) error
	Login(email, plainPassword string) (User, string, error)
}

type UserRepository interface {
	Insert(fullname, email, hashedPassword string, isVerified bool) error
	GetByEmail(email string) (User, error)
}
