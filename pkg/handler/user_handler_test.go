package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/anandawira/anandapay/pkg/usecase"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
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
}

func (ts *UserHandlerTestSuite) TestRegister() {
	ts.T().Run("It should return with StatusOK", func(t *testing.T) {
		ts.mockUsecase.On(
			"Register",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		form := url.Values{}
		form.Set("fullname", "test name")
		form.Set("email", "test@gmail.com")
		form.Set("password", "testpassword")

		req := httptest.NewRequest("POST", "/users", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		ts.handler.RegisterPost(c)
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)
	})

	ts.T().Run("It should return with StatusBadRequest on invalid input", func(t *testing.T) {
		req := httptest.NewRequest("POST", "/users", nil)
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		ts.handler.RegisterPost(c)
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})

	ts.T().Run("It should return with StatusBadRequest on duplicate email", func(t *testing.T) {
		const email string = "duplicate@gmail.com"

		ts.mockUsecase.On(
			"Register",
			mock.Anything,
			mock.AnythingOfType("string"),
			email,
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		form := url.Values{}
		form.Set("fullname", "test name")
		form.Set("email", email)
		form.Set("password", "testpassword")

		req := httptest.NewRequest("POST", "/users", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rec := httptest.NewRecorder()
		c, _ := gin.CreateTestContext(rec)
		c.Request = req

		ts.handler.RegisterPost(c)
		assert.Equal(t, http.StatusOK, rec.Result().StatusCode)

		ts.mockUsecase.On(
			"Register",
			mock.Anything,
			mock.AnythingOfType("string"),
			email,
			mock.AnythingOfType("string"),
		).Return(errors.New("Duplicate")).Once()

		req = httptest.NewRequest("POST", "/users", strings.NewReader(form.Encode()))
		req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
		rec = httptest.NewRecorder()
		c, _ = gin.CreateTestContext(rec)
		c.Request = req

		ts.handler.RegisterPost(c)
		assert.Equal(t, http.StatusBadRequest, rec.Result().StatusCode)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
