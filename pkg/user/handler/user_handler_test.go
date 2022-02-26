package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/anandawira/anandapay/domain"
	"github.com/anandawira/anandapay/pkg/user/usecase"
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
	gin.SetMode(gin.TestMode)
}

func createPostContext(body map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
	form := url.Values{}
	for key, val := range body {
		form.Set(key, val)
	}

	req := httptest.NewRequest("POST", "/users", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	return c, rec
}

func assertResponse(t testing.TB, code int, message string, recorder *httptest.ResponseRecorder) {
	t.Helper()

	want := gin.H{
		"message": message,
	}

	var got gin.H
	err := json.Unmarshal(recorder.Body.Bytes(), &got)
	if err != nil {
		t.Fatalf(err.Error())
	}

	assert.Equal(t, code, recorder.Code)
	assert.Equal(t, want, got)
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
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(nil).Once()

		c, rec := createPostContext(body)

		ts.handler.RegisterPost(c)

		assertResponse(t, http.StatusOK, "User registered to the database successfully.", rec)
	})

	ts.T().Run("It should return with StatusBadRequest on invalid input", func(t *testing.T) {
		c, rec := createPostContext(map[string]string{})

		ts.handler.RegisterPost(c)

		assertResponse(t, http.StatusBadRequest, domain.ErrParameterValidation.Error(), rec)
	})

	ts.T().Run("It should return with StatusBadRequest on duplicate email", func(t *testing.T) {
		ts.mockUsecase.On(
			"Register",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return(domain.ErrEmailUsed).Once()

		c, rec := createPostContext(body)

		ts.handler.RegisterPost(c)

		assertResponse(t, http.StatusBadRequest, domain.ErrEmailUsed.Error(), rec)
	})
}

func (ts *UserHandlerTestSuite) TestLogin() {
	body := map[string]string{
		"email":    "test@gmail.com",
		"password": "testpassword",
	}
	ts.T().Run("It should return with StatusOK and set cookie on login success", func(t *testing.T) {
		ts.mockUsecase.On(
			"Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return("token", nil).Once()

		c, rec := createPostContext(body)

		ts.handler.LoginPost(c)

		assert.NotEqual(t, "", rec.Header().Get("Set-Cookie"))
		assertResponse(t, http.StatusOK, "User logged in successfully.", rec)
	})

	ts.T().Run("It should return with StatusBadRequest and not set cookie on invalid input", func(t *testing.T) {
		c, rec := createPostContext(map[string]string{})

		ts.handler.LoginPost(c)

		assert.Equal(t, "", rec.Header().Get("Set-Cookie"))
		assertResponse(t, http.StatusBadRequest, domain.ErrParameterValidation.Error(), rec)
	})

	ts.T().Run("It should return with StatusBadRequest and not set cookie on wrong email or password", func(t *testing.T) {
		ts.mockUsecase.On(
			"Login",
			mock.Anything,
			mock.AnythingOfType("string"),
			mock.AnythingOfType("string"),
		).Return("", domain.ErrWrongEmailPass).Once()

		c, rec := createPostContext(body)

		ts.handler.LoginPost(c)

		assert.Equal(t, "", rec.Header().Get("Set-Cookie"))
		assertResponse(t, http.StatusBadRequest, domain.ErrWrongEmailPass.Error(), rec)
	})
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(UserHandlerTestSuite))
}
