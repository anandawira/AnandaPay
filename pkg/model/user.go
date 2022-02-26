package model

import (
	"context"

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
	Register(ctx context.Context, fullname, email, plainPassword string) error
	Login(ctx context.Context, email, plainPassword string) (token string, err error)
}

type UserRepository interface {
	Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error
	GetByEmail(ctx context.Context, email string) (user User, err error)
}
