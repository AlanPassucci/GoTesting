package tests

import (
	"bytes"
	"gotesting/04-TestsFuncionales/cmd/server"
	"gotesting/04-TestsFuncionales/prey"
	"gotesting/04-TestsFuncionales/shark"
	"gotesting/04-TestsFuncionales/simulator"
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {
	r := gin.Default()
	sim := simulator.NewCatchSimulator(35.4)

	whiteShark := shark.CreateWhiteShark(sim)
	tuna := prey.CreateTuna()

	handler := server.NewHandler(whiteShark, tuna)

	g := r.Group("/v1")

	g.PUT("/shark", handler.ConfigureShark())
	g.PUT("/prey", handler.ConfigurePrey())
	g.POST("/simulate", handler.SimulateHunt())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))

	req.Header.Add("Content-Type", "application/json")
	return req, httptest.NewRecorder()
}
