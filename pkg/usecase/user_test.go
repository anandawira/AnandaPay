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
}

func (m *MockUserRepo) Insert(ctx context.Context, fullname, email, hashedPassword string, isVerified bool) error {
	args := m.Called(ctx, fullname, email, hashedPassword, isVerified)
	return args.Error(0)
}

type UserUsecaseTestSuite struct {
	suite.Suite
	repo    *MockUserRepo
	usecase model.UserUsecase
}

func (ts *UserUsecaseTestSuite) SetupSuite() {
	ts.repo = new(MockUserRepo)
	ts.usecase = NewUserUsecase(ts.repo, 5)
}

func (ts *UserUsecaseTestSuite) TestRegister() {
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
			"useremail@gmail.com",
			"password",
		)

		ts.Assertions.NoError(err)
		ts.repo.AssertExpectations(t)
	})

	ts.T().Run("It should return false if email already exist.", func(t *testing.T) {
		const email string = "duplicate@gmail.com"
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
			"password2",
		)
		ts.Assertions.NoError(err)
		ts.repo.AssertExpectations(t)

		ts.repo.On(
			"Insert",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(errors.New("Duplicate email address")).Once()

		err = ts.usecase.Register(
			context.TODO(),
			"fullname2",
			email,
			"password2",
		)
		ts.Assertions.Error(err)
		ts.repo.AssertExpectations(t)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
