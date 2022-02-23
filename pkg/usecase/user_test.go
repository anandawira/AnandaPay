package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MockUserRepo struct {
	mock.Mock
	Emails map[string]bool
}

func (m *MockUserRepo) Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error {
	m.Called(ctx, fullname, email, hashedPassword, isVerified)
	if m.Emails[email] {
		return errors.New("Duplicate email address")
	}

	m.Emails[email] = true
	return nil
}

type UserUsecaseTestSuite struct {
	suite.Suite
	repo    *MockUserRepo
	usecase model.UserUsecase
}

func (ts *UserUsecaseTestSuite) SetupSuite() {
	ts.repo = &MockUserRepo{
		Emails: make(map[string]bool),
	}
	ts.usecase = NewUserUsecase(ts.repo, 5)
}

func (ts *UserUsecaseTestSuite) TestRegister() {
	const email string = "useremail@gmail.com"
	ts.T().Run("It should return true if user added to the database successfully.", func(t *testing.T) {

		ts.repo.On(
			"Insert",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(nil).Once()

		err := ts.usecase.Register(
			context.TODO(),
			"fullname1",
			email,
			"password",
		)

		ts.Assertions.NoError(err)
		ts.repo.AssertExpectations(t)
	})

	ts.T().Run("It should return false if email already exist.", func(t *testing.T) {
		ts.repo.On(
			"Insert",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(errors.New("Duplicate email address")).Once()

		err := ts.usecase.Register(
			context.TODO(),
			"fullname1",
			email,
			"password",
		)
		ts.Assertions.Error(err)
		ts.repo.AssertExpectations(t)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
