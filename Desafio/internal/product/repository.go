package product

import (
	"gotesting/Desafio/internal/domain"

	"github.com/stretchr/testify/mock"
)

type Repository interface {
	GetAllBySeller(sellerID string) ([]domain.Product, error)
}

type repository struct {
	productsList []domain.Product
}

func NewRepository(productsList []domain.Product) Repository {
	return &repository{
		productsList: productsList,
	}
}

func (r *repository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	products := make([]domain.Product, 0)

	for _, product := range r.productsList {
		if product.SellerID == sellerID {
			products = append(products, product)
		}
	}

	return products, nil
}

type StubRepository struct {
	Products []domain.Product
	Error    error
}

func (sr *StubRepository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	return sr.Products, sr.Error
}

type MockRepository struct {
	mock.Mock
}

func (mr *MockRepository) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	args := mr.Called(sellerID)
	return args.Get(0).([]domain.Product), args.Error(1)
}
