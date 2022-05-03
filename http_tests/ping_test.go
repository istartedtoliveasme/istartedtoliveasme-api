package http_tests

import (
	routers2 "api/configs/routers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPingRoute(t *testing.T) {
	router := gin.Default()
	routers2.GetV1Routers(router)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, routers2.GetURLPath(routers2.Version1)+routers2.GetURLPath(routers2.Ping), nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "pong")
}
