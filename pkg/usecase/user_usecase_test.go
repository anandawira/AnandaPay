package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/anandawira/anandapay/pkg/model"
	"github.com/anandawira/anandapay/pkg/repo"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo *repo.MockUserRepo
	usecase  model.UserUsecase
}

func (ts *UserUsecaseTestSuite) SetupSuite() {
	ts.mockRepo = new(repo.MockUserRepo)
	ts.usecase = NewUserUsecase(ts.mockRepo, 5)
}

func (ts *UserUsecaseTestSuite) TestRegister() {
	ts.T().Run("It should return true if user added to the database successfully.", func(t *testing.T) {

		ts.mockRepo.On(
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
		ts.mockRepo.AssertExpectations(t)
	})

	ts.T().Run("It should return false if email already exist.", func(t *testing.T) {
		const email string = "duplicate@gmail.com"

		ts.mockRepo.On(
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
		ts.mockRepo.AssertExpectations(t)

		ts.mockRepo.On(
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
		ts.mockRepo.AssertExpectations(t)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
