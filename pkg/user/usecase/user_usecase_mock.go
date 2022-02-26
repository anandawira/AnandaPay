package usecase

import (
	"context"

	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(ctx context.Context, fullname, email, plainPassword string) error {
	args := m.Called(ctx, fullname, email, plainPassword)
	return args.Error(0)
}

func (m *MockUserUsecase) Login(ctx context.Context, email string, plainPassword string) (domain.User, string, error) {
	args := m.Called(ctx, email, plainPassword)
	return args.Get(0).(domain.User), args.String(1), args.Error(2)
}
