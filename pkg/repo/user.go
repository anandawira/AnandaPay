package repo

import (
	"context"

	"github.com/anandawira/anandapay/pkg/model"
	"gorm.io/gorm"
)

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{DB: db}
}

func (m *userRepository) Insert(ctx context.Context, fullname, email, hashedPassword string) error {
	user := model.User{FullName: fullname, Email: email, HashedPassword: hashedPassword, IsVerified: false}
	result := m.DB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
