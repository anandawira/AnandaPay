package usecase

import (
	"errors"
	"testing"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/repo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
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
	ts.usecase = NewUserUsecase(ts.mockRepo)
}

func (ts *UserUsecaseTestSuite) TestRegister() {
	ts.T().Run("It should return true if user added to the database successfully.", func(t *testing.T) {
		ts.mockRepo.On(
			"Insert",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(nil).Once()

		err := ts.usecase.Register(
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
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("bool"),
		).Return(errors.New("Duplicate email address")).Once()

		err := ts.usecase.Register(
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
	ts.T().Run("It should return user and token if email and password match", func(t *testing.T) {
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), 0)
		require.NoError(t, err)

		user := domain.User{
			FullName:       "user name 1",
			Email:          "email@gmail.com",
			HashedPassword: string(hashedPassword),
			IsVerified:     false,
		}

		wallet := domain.Wallet{
			ID:     "wallet id",
			UserID: user.ID,
		}

		ts.mockRepo.On(
			"GetByEmail",
			mock.AnythingOfType("string"),
		).Return(user, wallet, nil).Once()

		userLogin, userWallet, token, err := ts.usecase.Login("email", plainPassword)
		require.NoError(t, err)
		assert.Equal(t, user, userLogin)
		assert.Equal(t, wallet, userWallet)
		assert.Equal(t, userLogin.ID, userWallet.UserID)
		assert.NotEqual(t, "", token)
	})

	ts.T().Run("It should return error if email not found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetByEmail",
			mock.AnythingOfType("string"),
		).Return(domain.User{}, domain.Wallet{}, nil).Once()

		_, _, _, err := ts.usecase.Login("email", plainPassword)
		assert.Error(t, err)
	})

	ts.T().Run("It should return error if email and password doesn't match", func(t *testing.T) {

		user := domain.User{
			FullName:       "user name 1",
			Email:          "email@gmail.com",
			HashedPassword: "randomHashedPassword",
			IsVerified:     false,
		}

		ts.mockRepo.On(
			"GetByEmail",
			mock.AnythingOfType("string"),
		).Return(user, domain.Wallet{}, nil).Once()

		_, _, _, err := ts.usecase.Login("email", "anotherPassword")
		assert.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
