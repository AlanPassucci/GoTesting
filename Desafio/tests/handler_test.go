package tests

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandler_GetProductsBySellerID(t *testing.T) {
	t.Run("should return a seller products list by the given seller_id", func(t *testing.T) {
		// Arrange
		server := createTestServer()
		givenSellerID := "FEX112AC"

		expectedStatusCode := 200
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `[{"id":"mock","seller_id":"FEX112AC","description":"generic product","price":123.55}]`

		req, res := createTestRequest("GET", "/api/v1/products?seller_id="+givenSellerID, "")

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should return an empty seller products list because that seller does not have products", func(t *testing.T) {
		// Arrange
		server := createTestServer()
		givenSellerID := "ABC123DE"

		expectedStatusCode := 200
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `[]`

		req, res := createTestRequest("GET", "/api/v1/products?seller_id="+givenSellerID, "")

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})

	t.Run("should return a 400 bad request when the given seller_id is empty", func(t *testing.T) {
		// Arrange
		server := createTestServer()
		givenSellerID := ""

		expectedStatusCode := 400
		expectedHeaders := http.Header{
			"Content-Type": []string{"application/json; charset=utf-8"},
		}
		expectedBody := `{"error":"seller_id query param is required"}`

		req, res := createTestRequest("GET", "/api/v1/products?seller_id="+givenSellerID, "")

		// Act
		server.ServeHTTP(res, req)

		// Assert
		assert.Equal(t, expectedStatusCode, res.Code)
		assert.Equal(t, expectedHeaders, res.Header())
		assert.JSONEq(t, expectedBody, res.Body.String())
	})
}
