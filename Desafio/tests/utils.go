package tests

import (
	"bytes"
	"gotesting/Desafio/cmd/router"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createTestServer() *gin.Engine {
	server := gin.Default()
	router.MapRoutes(server)

	return server
}

func createTestRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
