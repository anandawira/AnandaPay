package repo

import (
	"github.com/anandawira/anandapay/domain"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return &userRepository{db: db}
}

func (m *userRepository) Insert(fullname, email, hashedPassword string, isVerified bool) error {
	err := m.db.Transaction(func(tx *gorm.DB) error {
		user := domain.User{
			FullName:       fullname,
			Email:          email,
			HashedPassword: hashedPassword,
			IsVerified:     isVerified,
		}
		result := m.db.Create(&user)
		if result.Error != nil {
			return result.Error
		}

		wallet := domain.Wallet{
			ID:      uuid.NewString(),
			UserID:  user.ID,
			Balance: 0,
		}
		result = m.db.Create(&wallet)
		if result.Error != nil {
			return result.Error
		}

		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (m *userRepository) GetByEmail(email string) (user domain.User, wallet domain.Wallet, err error) {

	result := m.db.Select("id", "full_name", "email", "hashed_password").Where("email = ?", email).First(&user)
	if result.Error != nil {
		return domain.User{}, domain.Wallet{}, domain.ErrEmailNotFound
	}

	result = m.db.Select("id", "user_id").Where("user_id = ?", user.ID).First(&wallet)
	if result.Error != nil {
		return domain.User{}, domain.Wallet{}, domain.ErrWalletNotFound
	}

	return user, wallet, nil
}
