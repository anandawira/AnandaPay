package repo

import (
	"context"

	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error {
	args := m.Called(ctx, fullname, email, hashedPassword, isVerified)
	return args.Error(0)
}

func (m *MockUserRepo) GetByEmail(ctx context.Context, email string) (user domain.User, err error) {
	args := m.Called(ctx, email)
	return args.Get(0).(domain.User), args.Error(1)
}
