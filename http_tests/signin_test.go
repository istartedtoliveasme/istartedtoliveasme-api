package http_tests

import (
	"api/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestSignInShouldNotAuthorizedWhenEmptyBody(t *testing.T) {
	router := gin.Default()
	routers.GetV1Routers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, routers.GetURLPath(routers.Version1)+routers.GetURLPath(routers.SignIn), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Contains(t, w.Body.String(), "error")
}
