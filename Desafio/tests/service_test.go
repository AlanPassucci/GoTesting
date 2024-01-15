package tests

import (
	"gotesting/Desafio/internal/domain"
	"gotesting/Desafio/internal/product"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductsService_GetAllBySeller(t *testing.T) {
	t.Run("should return a list of products by the given seller_id", func(t *testing.T) {
		// Arrange
		givenSellerID := "FEX112AC"
		expectedProducts := []domain.Product{
			{ID: "mock", SellerID: "FEX112AC", Description: "generic product", Price: 123.55},
		}

		stubRepo := &product.StubRepository{
			Products: expectedProducts,
			Error:    nil,
		}
		mockRepo := &product.MockRepository{}
		service := product.NewService(mockRepo)

		mockRepo.On("GetAllBySeller", givenSellerID).Return(stubRepo.Products, stubRepo.Error)

		// Act
		obtainedProducts, err := service.GetAllBySeller(givenSellerID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, obtainedProducts)
		assert.True(t, mockRepo.AssertExpectations(t))
	})

	t.Run("should return an empty list of products because the given seller does not have any", func(t *testing.T) {
		// Arrange
		givenSellerID := "ABC123DE"
		expectedProducts := make([]domain.Product, 0)

		stubRepo := &product.StubRepository{
			Products: expectedProducts,
			Error:    nil,
		}
		mockRepo := &product.MockRepository{}
		service := product.NewService(mockRepo)

		mockRepo.On("GetAllBySeller", givenSellerID).Return(stubRepo.Products, stubRepo.Error)

		// Act
		obtainedProducts, err := service.GetAllBySeller(givenSellerID)

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedProducts, obtainedProducts)
		assert.True(t, mockRepo.AssertExpectations(t))
	})
}
