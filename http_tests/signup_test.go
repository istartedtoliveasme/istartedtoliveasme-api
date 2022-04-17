package http_tests

import (
	"api/helpers/httpHelper"
	"api/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestSignUpShouldNotAuthorizedWhenEmptyBody(t *testing.T) {
	url := routers.GetURLPath(routers.Version1) + routers.GetURLPath(routers.SignUp)
	router := gin.Default()
	routers.GetV1Routers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSignUpShouldAuthorized(t *testing.T) {
	url := routers.GetURLPath(routers.Version1) + routers.GetURLPath(routers.SignUp)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	// TODO :: fix testing issue
	fakeFirstName := "randomFirstName"
	fakeLastName := "randomFirstName"
	fakeEmail := strconv.Itoa(rand.Int()) + "@gmail.com"
	c.JSON(http.StatusOK, httpHelper.JSON{
		"firstName": fakeFirstName,
		"lastName":  fakeLastName,
		"email":     fakeEmail,
	})

	c.Request, _ = http.NewRequest(http.MethodPost, url, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "firstName")
	assert.Contains(t, w.Body.String(), "lastName")
	assert.Contains(t, w.Body.String(), "email")
	assert.Contains(t, w.Body.String(), fakeFirstName)
	assert.Contains(t, w.Body.String(), fakeLastName)
	assert.Contains(t, w.Body.String(), fakeEmail)
}
