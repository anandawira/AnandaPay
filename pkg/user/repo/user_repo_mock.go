package repo

import (
	"github.com/anandawira/anandapay/domain"
	"github.com/stretchr/testify/mock"
)

type MockUserRepo struct {
	mock.Mock
}

func (m *MockUserRepo) Insert(fullname, email, hashedPassword string, isVerified bool) error {
	args := m.Called(fullname, email, hashedPassword, isVerified)
	return args.Error(0)
}

func (m *MockUserRepo) GetByEmail(email string) (user domain.User, err error) {
	args := m.Called(email)
	return args.Get(0).(domain.User), args.Error(1)
}
