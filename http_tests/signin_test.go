package http_tests

import (
	routers2 "api/configs/routers"
	"api/constants"
	helperTypes "api/helpers/typings"
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSignInShouldNotAuthorizedWhenEmptyBody(t *testing.T) {
	url := routers2.UrlPath(routers2.Version1)
	url.Slugs(routers2.SignIn)
	router := gin.Default()
	routers2.GetV1Routers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, string(url), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSignInShouldThrowPasswordNotMatch(t *testing.T) {
	// GIVEN the initial dependency is declared
	url := routers2.UrlPath(routers2.Version1)
	url.Slugs(routers2.SignIn)
	router := gin.Default()
	routers2.GetV1Routers(router)
	w := httptest.NewRecorder()
	// AND I have the username
	fakeUsername := "istartedtoliveasme"
	// AND I have a non-existing password
	fakePassword := strconv.Itoa(rand.Int())
	// AND I have the serialized json format of the API body
	jsonBody := helperTypes.Json[helperTypes.JsonPayload]{
		Payload: helperTypes.JsonPayload{
			"username": fakeUsername,
			"password": fakePassword,
		},
	}

	jsonByte, err := json.Marshal(jsonBody)
	// AND Should throw error when fail to parse struct
	if err != nil {
		t.Error(err)
	}

	jsonBuffer := bytes.NewBuffer(jsonByte)

	// WHEN I invoke the http request
	req, _ := http.NewRequest(http.MethodPost, string(url), jsonBuffer)

	router.ServeHTTP(w, req)

	// THEN I should be able to verify the API responses status code
	assert.Equal(t, http.StatusBadRequest, w.Code)
	// AND Should be able to see the error list
	assert.Contains(t, w.Body.String(), "errors")
	// AND Should be able to see the error message
	assert.Contains(t, w.Body.String(), constants.MismatchPassword)
}

func TestSignInShouldAuthorized(t *testing.T) {
	url := routers2.UrlPath(routers2.Version1)
	url.Slugs(routers2.SignIn)
	router := gin.Default()
	routers2.GetV1Routers(router)

	w := httptest.NewRecorder()

	fakeUsername := "istartedtoliveasme"
	fakePassword := strconv.Itoa(rand.Int())

	jsonBody := helperTypes.Json[helperTypes.JsonPayload]{
		Payload: helperTypes.JsonPayload{
			"username": fakeUsername,
			"password": "123456",
		},
	}

	jsonByte, err := json.Marshal(jsonBody)

	if err != nil {
		t.Error(err)
	}

	jsonBuffer := bytes.NewBuffer(jsonByte)

	req, _ := http.NewRequest(http.MethodPost, string(url), jsonBuffer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Id")
	assert.Contains(t, w.Body.String(), "FirstName")
	assert.Contains(t, w.Body.String(), "LastName")
	assert.Contains(t, w.Body.String(), "Email")
	assert.Contains(t, w.Body.String(), fakeUsername+"@gmail.com")
	assert.NotContains(t, w.Body.String(), fakePassword)
}
