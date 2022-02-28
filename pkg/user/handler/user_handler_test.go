package handler

import (
	"net/http"
	"testing"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/helper"
	"github.com/anandawira/anandapay/pkg/user/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserHandlerTestSuite struct {
	suite.Suite
	mockUsecase *usecase.MockUserUsecase
	handler     *UserHandler
}

func (ts *UserHandlerTestSuite) SetupSuite() {
	ts.mockUsecase = new(usecase.MockUserUsecase)
	ts.handler = &UserHandler{
		userUsecase: ts.mockUsecase,
	}
	gin.SetMode(gin.TestMode)
}

func (ts *UserHandlerTestSuite) TestRegister() {
	body := map[string]string{
		"fullname": "testname",
		"email":    "test@gmail.com",
		"password": "testpassword",
	}
	ts.T().Run("It should return with StatusOK", func(t *testing.T) {
		ts.mockUsecase.On(
			"Register",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		c, rec := helper.CreatePostContext(body)

		ts.handler.RegisterPost(c)

		helper.AssertResponse(t, http.StatusOK, gin.H{"message": "User registered to the database successfully."}, rec)
	})

	ts.T().Run("It should return with StatusBadRequest on invalid input", func(t *testing.T) {
		c, rec := helper.CreatePostContext(map[string]string{})

		ts.handler.RegisterPost(c)

		helper.AssertResponse(t, http.StatusBadRequest, gin.H{"message": domain.ErrParameterValidation.Error()}, rec)
	})

	ts.T().Run("It should return with StatusBadRequest on duplicate email", func(t *testing.T) {
		ts.mockUsecase.On(
			"Register",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(domain.ErrEmailUsed).Once()

		c, rec := helper.CreatePostContext(body)

		ts.handler.RegisterPost(c)

		helper.AssertResponse(t, http.StatusBadRequest, gin.H{"message": domain.ErrEmailUsed.Error()}, rec)
	})
}

func (ts *UserHandlerTestSuite) TestLogin() {
	user := domain.User{
		FullName: "Full Name",
		Email:    "test@gmail.com",
	}

	wallet := domain.Wallet{
		ID:     "wallet id",
		UserID: user.ID,
	}

	body := map[string]string{
		"email":    user.Email,
		"password": "testpassword",
	}

	ts.T().Run("It should return with StatusOK and data on login success", func(t *testing.T) {
		ts.mockUsecase.On(
			"Login",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(user, wallet, "token", nil).Once()

		c, rec := helper.CreatePostContext(body)

		ts.handler.LoginPost(c)

		body := gin.H{
			"message": "User logged in successfully.",
			"data": LoginResponseData{
				UserID:      0,
				WalletID:    wallet.ID,
				Fullname:    user.FullName,
				Email:       user.Email,
				AccessToken: "token",
			},
		}
		helper.AssertResponse(t, http.StatusOK, body, rec)
	})

	ts.T().Run("It should return with StatusBadRequest on invalid input", func(t *testing.T) {
		c, rec := helper.CreatePostContext(map[string]string{})

		ts.handler.LoginPost(c)

		helper.AssertResponse(t, http.StatusBadRequest, gin.H{"message": domain.ErrParameterValidation.Error()}, rec)
	})

	ts.T().Run("It should return with StatusBadRequest on wrong email or password", func(t *testing.T) {
		ts.mockUsecase.On(
			"Login",
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(user, wallet, "", domain.ErrWrongEmailPass).Once()

		c, rec := helper.CreatePostContext(body)

		ts.handler.LoginPost(c)

		helper.AssertResponse(t, http.StatusBadRequest, gin.H{"message": domain.ErrWrongEmailPass.Error()}, rec)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
