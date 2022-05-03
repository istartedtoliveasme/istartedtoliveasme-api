package http_tests

import (
	routers2 "api/configs/routers"
	"api/helpers/httpHelper"
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

func TestSignUpShouldNotAuthorizedWhenEmptyBody(t *testing.T) {
	url := routers2.GetURLPath(routers2.Version1) + routers2.GetURLPath(routers2.SignUp)
	router := gin.Default()
	routers2.GetV1Routers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSignUpShouldAuthorized(t *testing.T) {
	url := routers2.GetURLPath(routers2.Version1) + routers2.GetURLPath(routers2.SignUp)
	router := gin.Default()
	routers2.GetV1Routers(router)

	w := httptest.NewRecorder()

	fakeFirstName := "randomFirstName"
	fakeLastName := "randomFirstName"
	fakeEmail := strconv.Itoa(rand.Int()) + "@gmail.com"

	jsonBody := httpHelper.JSON{
		"firstName": fakeFirstName,
		"lastName":  fakeLastName,
		"email":     fakeEmail,
	}

	jsonByte, err := json.Marshal(jsonBody)

	if err != nil {
		t.Error(err)
	}

	jsonBuffer := bytes.NewBuffer(jsonByte)

	req, _ := http.NewRequest(http.MethodPost, url, jsonBuffer)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "data")
	assert.Contains(t, w.Body.String(), "FirstName")
	assert.Contains(t, w.Body.String(), fakeFirstName)
	assert.Contains(t, w.Body.String(), "LastName")
	assert.Contains(t, w.Body.String(), fakeLastName)
	assert.Contains(t, w.Body.String(), "Email")
	assert.Contains(t, w.Body.String(), fakeEmail)
}
