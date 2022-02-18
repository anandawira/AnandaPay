package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `json:"fullname"`
	Email          string `json:"email" gorm:"unique"`
	HashedPassword string `json:"password"`
	IsVerified     bool   `json:"isVerified"`
}

type UserUsecase interface {
	Register(ctx context.Context, fullname, email, plainPassword string) error
}

type UserRepository interface {
	Insert(ctx context.Context, fullname, email, hashedPassword string) error
}
