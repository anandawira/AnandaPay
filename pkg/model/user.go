package model

import (
	"context"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FullName       string `json:"fullname"`
	Email          string `json:"email"`
	HashedPassword string `json:"password"`
	isVerified     string `json:"isVerified"`
}

type UserUsecase interface {
	Register(ctx context.Context) error
}

type UserRepository interface {
	Insert(ctx context.Context, fullname, email, hashedPassword string) error
}
