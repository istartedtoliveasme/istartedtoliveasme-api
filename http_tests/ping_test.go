package http_tests

import (
	"api/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	defaultRoute := gin.Default()
	router := routers.GetV1Routers(defaultRoute)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/v1/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "pong")
}
