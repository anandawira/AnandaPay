package usecase

import (
	"errors"
	"strconv"
	"testing"
	"time"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/repo"
	"github.com/golang-jwt/jwt"
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
	ts.usecase = NewUserUsecase(ts.mockRepo, 5)
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

		ts.mockRepo.On(
			"GetByEmail",
			mock.AnythingOfType("string"),
		).Return(user, nil).Once()

		userLogin, token, err := ts.usecase.Login("email", plainPassword)
		require.NoError(t, err)
		assert.Equal(t, user, userLogin)
		assert.NotEqual(t, "", token)
	})

	ts.T().Run("It should return error if email not found", func(t *testing.T) {
		ts.mockRepo.On(
			"GetByEmail",
			mock.AnythingOfType("string"),
		).Return(domain.User{}, nil).Once()

		_, _, err := ts.usecase.Login("email", plainPassword)
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
		).Return(user, nil).Once()

		_, _, err := ts.usecase.Login("email", "anotherPassword")
		assert.Error(t, err)
	})
}

func (ts *UserUsecaseTestSuite) TestVerifyToken() {
	// Hardcode, later change to env
	var secretKey string = "secret"
	const id int = 1

	ts.T().Run("It should return userId on valid token", func(t *testing.T) {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(id),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		})

		validToken, err := claims.SignedString([]byte(secretKey))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		userId, err := ts.usecase.VerifyToken(validToken)
		require.NoError(t, err)
		assert.Equal(t, id, userId)
	})
	ts.T().Run("It should return error on invalid token", func(t *testing.T) {
		_, err := ts.usecase.VerifyToken("invalid token")
		require.Error(t, err)
	})
	ts.T().Run("It should return error on token signed with different secret key", func(t *testing.T) {
		claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
			Issuer:    strconv.Itoa(id),
			ExpiresAt: time.Now().Add(12 * time.Hour).Unix(),
		})

		token, err := claims.SignedString([]byte("other secret key"))
		if err != nil {
			t.Fatal("JWT token generation failed.", err.Error())
		}

		_, err = ts.usecase.VerifyToken(token)
		require.Error(t, err)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}
