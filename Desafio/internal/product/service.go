package product

import (
	"fmt"
	"gotesting/Desafio/internal/domain"
)

type Service interface {
	GetAllBySeller(sellerID string) ([]domain.Product, error)
}

type service struct {
	rp Repository
}

func NewService(rp Repository) Service {
	return &service{
		rp: rp,
	}
}

func (sv *service) GetAllBySeller(sellerID string) ([]domain.Product, error) {
	data, err := sv.rp.GetAllBySeller(sellerID)
	if err != nil {
		fmt.Println("error in repository", err.Error(), "sellerId:", sellerID)
		return nil, err
	}

	return data, err
}
