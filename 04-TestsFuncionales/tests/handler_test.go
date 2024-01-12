package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_ConfigureShark(t *testing.T) {
	t.Run("should return success true because shark config has valid field types", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBody := `
			{
				"x_position": 50.0,
				"y_position": 50.0,
				"speed": 11.2
			}
		`

		expectedStatusCode := http.StatusOK
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{"success": true}`

		req, res := createRequestTest(http.MethodPut, "/v1/shark", reqBody)

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should return success false because shark config has invalid field types", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBody := `
			{
				"x_position": "50.0",
				"y_position": "50.0",
				"speed": "11.2"
			}
		`

		expectedStatusCode := http.StatusBadRequest
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{"success": false}`

		req, res := createRequestTest(http.MethodPut, "/v1/shark", reqBody)

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestHandler_ConfigurePrey(t *testing.T) {
	t.Run("should return success true because prey config has valid field types", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBody := `{"speed": 21.0}`

		expectedStatusCode := http.StatusOK
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{"success": true}`

		req, res := createRequestTest(http.MethodPut, "/v1/prey", reqBody)

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should return success false because prey config has invalid field types", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBody := `{"speed": "21.0"}`

		expectedStatusCode := http.StatusBadRequest
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{"success": false}`

		req, res := createRequestTest(http.MethodPut, "/v1/prey", reqBody)

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})
}

func TestHandler_SimulateHunt(t *testing.T) {
	t.Run("should return success false because prey is faster than shark", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBodyConfigPrey := `{"speed": 21.0}`
		reqBodyConfigShark := `
			{
				"x_position": 50.0,
				"y_position": 50.0,
				"speed": 11.2
			}
		`

		expectedStatusCode := http.StatusBadRequest
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `
			{
				"success": false, 
				"message": "could not catch it", 
				"time": 0
			}
		`

		reqConfigPrey, resConfigPrey := createRequestTest(http.MethodPut, "/v1/prey", reqBodyConfigPrey)
		reqConfigShark, resConfigShark := createRequestTest(http.MethodPut, "/v1/shark", reqBodyConfigShark)
		reqSimulateHunt, resSimulateHunt := createRequestTest(http.MethodPost, "/v1/simulate", "")

		// Act
		server.ServeHTTP(resConfigPrey, reqConfigPrey)
		server.ServeHTTP(resConfigShark, reqConfigShark)
		server.ServeHTTP(resSimulateHunt, reqSimulateHunt)

		// Assert
		assert.Equal(t, expectedStatusCode, resSimulateHunt.Code)
		assert.Equal(t, expectedHeaders, resSimulateHunt.Header())
		assert.JSONEq(t, expectedBody, resSimulateHunt.Body.String())
	})

	t.Run("should return success false because shark is faster than prey but the distance is longer", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBodyConfigPrey := `{"speed": 10.0}`
		reqBodyConfigShark := `
			{
				"x_position": 50.0,
				"y_position": 50.0,
				"speed": 11.2
			}
		`

		expectedStatusCode := http.StatusBadRequest
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `
			{
				"success": false, 
				"message": "could not catch it", 
				"time": 0
			}
		`

		reqConfigPrey, resConfigPrey := createRequestTest(http.MethodPut, "/v1/prey", reqBodyConfigPrey)
		reqConfigShark, resConfigShark := createRequestTest(http.MethodPut, "/v1/shark", reqBodyConfigShark)
		reqSimulateHunt, resSimulateHunt := createRequestTest(http.MethodPost, "/v1/simulate", "")

		// Act
		server.ServeHTTP(resConfigPrey, reqConfigPrey)
		server.ServeHTTP(resConfigShark, reqConfigShark)
		server.ServeHTTP(resSimulateHunt, reqSimulateHunt)

		// Assert
		assert.Equal(t, expectedStatusCode, resSimulateHunt.Code)
		assert.Equal(t, expectedHeaders, resSimulateHunt.Header())
		assert.JSONEq(t, expectedBody, resSimulateHunt.Body.String())
	})

	t.Run("should return success true because shark hunt the prey in 24 seconds", func(t *testing.T) {
		// Arrange
		server := createServer()
		reqBodyConfigPrey := `{"speed": 8.254}`
		reqBodyConfigShark := `
			{
				"x_position": 50.0,
				"y_position": 50.0,
				"speed": 11.2
			}
		`

		expectedStatusCode := http.StatusOK
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `
			{
				"success": true, 
				"message": "", 
				"time": 24
			}
		`

		reqConfigPrey, resConfigPrey := createRequestTest(http.MethodPut, "/v1/prey", reqBodyConfigPrey)
		reqConfigShark, resConfigShark := createRequestTest(http.MethodPut, "/v1/shark", reqBodyConfigShark)
		reqSimulateHunt, resSimulateHunt := createRequestTest(http.MethodPost, "/v1/simulate", "")

		// Act
		server.ServeHTTP(resConfigPrey, reqConfigPrey)
		server.ServeHTTP(resConfigShark, reqConfigShark)
		server.ServeHTTP(resSimulateHunt, reqSimulateHunt)

		// Assert
		assert.Equal(t, expectedStatusCode, resSimulateHunt.Code)
		assert.Equal(t, expectedHeaders, resSimulateHunt.Header())
		assert.JSONEq(t, expectedBody, resSimulateHunt.Body.String())
	})
}
