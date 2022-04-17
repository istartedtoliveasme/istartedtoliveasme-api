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

func TestSignInShouldNotAuthorizedWhenEmptyBody(t *testing.T) {
	url := routers.GetURLPath(routers.Version1) + routers.GetURLPath(routers.SignIn)
	router := gin.Default()
	routers.GetV1Routers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}

func TestSignInShouldAuthorized(t *testing.T) {
	url := routers.GetURLPath(routers.Version1) + routers.GetURLPath(routers.SignIn)
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	fakeUsername := "istartedtoliveasme"
	fakePassword := strconv.Itoa(rand.Int())
	c.JSON(http.StatusOK, httpHelper.JSON{
		"username": fakeUsername,
		"password": fakePassword,
	})

	c.Request, _ = http.NewRequest(http.MethodPost, url, nil)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "username")
	assert.Contains(t, w.Body.String(), "password")
	assert.Contains(t, w.Body.String(), fakeUsername)
	assert.Contains(t, w.Body.String(), fakePassword)
}
