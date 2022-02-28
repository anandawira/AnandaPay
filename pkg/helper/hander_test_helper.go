package helper

import (
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func CreatePostContext(body map[string]string) (*gin.Context, *httptest.ResponseRecorder) {
	form := url.Values{}
	for key, val := range body {
		form.Set(key, val)
	}

	req := httptest.NewRequest("POST", "/", strings.NewReader(form.Encode()))
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	return c, rec
}

func CreateGetContext() (*gin.Context, *httptest.ResponseRecorder) {
	req := httptest.NewRequest("GET", "/", nil)
	rec := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(rec)
	c.Request = req
	return c, rec
}

func AssertResponse(t testing.TB, code int, body gin.H, recorder *httptest.ResponseRecorder) {
	t.Helper()

	want, err := json.Marshal(body)
	if err != nil {
		t.Fatalf(err.Error())
	}

	got := recorder.Body.Bytes()

	assert.Equal(t, code, recorder.Code, "http status code not equal")
	assert.Equal(t, want, got, "response body not equal")
}
