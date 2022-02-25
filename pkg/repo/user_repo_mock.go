package repo

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error {
	args := m.Called(ctx, fullname, email, hashedPassword, isVerified)
	return args.Error(0)
}
