package repo

import (
	"context"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error {
	args := m.Called(ctx, fullname, email, hashedPassword, isVerified)
	return args.Error(0)
}

func (m *MockUserRepo) GetOne(ctx context.Context, email string, hashedPassword string) (user model.User, err error) {
	args := m.Called(ctx, email, hashedPassword)
	return args.Get(0).(model.User), args.Error(1)
}
