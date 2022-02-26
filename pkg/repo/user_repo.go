package repo

import (
	"context"

	"github.com/anandawira/anandapay/pkg/model"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) model.UserRepository {
	return &userRepository{db: db}
}

func (m *userRepository) Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error {
	user := model.User{FullName: fullname, Email: email, HashedPassword: hashedPassword, IsVerified: isVerified}
	result := m.db.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (m *userRepository) GetByEmail(ctx context.Context, email string) (user model.User, err error) {
	result := m.db.Where("email = ?", email).First(&user)
	err = result.Error
	return user, err
}
