package usecase

import (
	"context"
	"time"

	"github.com/anandawira/anandapay/pkg/model"
	"golang.org/x/crypto/bcrypt"
)

type userUsecase struct {
	userRepo       model.UserRepository
	contextTimeout time.Duration
}

func NewUserUsecase(repo model.UserRepository, timeout time.Duration) model.UserUsecase {
	return &userUsecase{userRepo: repo, contextTimeout: timeout}
}

func (m *userUsecase) Register(c context.Context, fullname, email, plainPassword string) error {
	ctx, cancel := context.WithTimeout(c, m.contextTimeout)
	defer cancel()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 14)
	if err != nil {
		return err
	}

	err = m.userRepo.Insert(ctx, fullname, email, string(hashedPassword), false)
	if err != nil {
		return err
	}

	return nil
}
