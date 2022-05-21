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
	basePath := routers2.UrlPath(routers2.Version1)
	url := basePath.Slugs(routers2.Ping)
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "message")
	assert.Contains(t, w.Body.String(), "pong")
}
