package usecase

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserUsecase struct {
	mock.Mock
}

func (m *MockUserUsecase) Register(ctx context.Context, fullname, email, plainPassword string) error {
	args := m.Called(ctx, fullname, email, plainPassword)
	return args.Error(0)
}
