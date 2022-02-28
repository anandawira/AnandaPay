package usecase

import (
	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(fullname, email, plainPassword string) error {
	args := m.Called(fullname, email, plainPassword)
	return args.Error(0)
}

func (m *MockUserUsecase) Login(email string, plainPassword string) (domain.User, domain.Wallet, string, error) {
	args := m.Called(email, plainPassword)
	return args.Get(0).(domain.User), args.Get(1).(domain.Wallet), args.String(2), args.Error(3)
}
