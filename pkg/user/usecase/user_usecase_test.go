package usecase

import (
	"context"
	"errors"
	"log"
	"testing"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecaseTestSuite struct {
	suite.Suite
	mockRepo *repo.MockUserRepo
	usecase  domain.UserUsecase
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

		assert.NoError(t, err)
		ts.mockRepo.AssertExpectations(t)
	})

	ts.T().Run("It should return false if email already exist.", func(t *testing.T) {
		ts.mockRepo.On(
			"Insert",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(errors.New("Duplicate email address")).Once()

		err := ts.usecase.Register(
			context.TODO(),
			"fullname2",
			"duplicate@gmail.com",
			"password2",
		)
		assert.Error(t, err)
		ts.mockRepo.AssertExpectations(t)
	})
}

func (ts *UserUsecaseTestSuite) TestLogin() {
	const plainPassword string = "plainPassword"
	ts.T().Run("It should return token if email and password math", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 0)
		if err != nil {
			log.Fatal("Password hashing error", err.Error())
		}

		user := domain.User{
			FullName:       "user name 1",
			Email:          "email@gmail.com",
			HashedPassword: string(hashedPassword),
			IsVerified:     false,
		}

		ts.mockRepo.On(
			"GetByEmail",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(user, nil).Once()

		_, err = ts.usecase.Login(context.TODO(), "email", plainPassword)
		assert.NoError(t, err)
	})

	ts.T().Run("It should return error if email not found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetByEmail",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(domain.User{}, nil).Once()

		_, err := ts.usecase.Login(context.TODO(), "email", plainPassword)
		assert.Error(t, err)
	})

	ts.T().Run("It should return error if email and password doesn't match", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 0)
		if err != nil {
			log.Fatal("Password hashing error", err.Error())
		}

		user := domain.User{
			FullName:       "user name 1",
			Email:          "email@gmail.com",
			HashedPassword: string(hashedPassword),
			IsVerified:     false,
		}

		ts.mockRepo.On(
			"GetByEmail",
			mock.Anything,
			mock.AnythingOfType("string"),
		).Return(user, nil).Once()

		_, err = ts.usecase.Login(context.TODO(), "email", "anotherPassword")
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
