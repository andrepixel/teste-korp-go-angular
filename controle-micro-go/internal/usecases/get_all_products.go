package usecases

import (
	"controle-micro-go/internal/entities"
	"controle-micro-go/internal/interfaces"
)

type GetAllProductsUseCase struct {
	repo interfaces.ProductRepositoryInterface
}

func NewGetAllProductsUseCase(repo interfaces.ProductRepositoryInterface) *GetAllProductsUseCase {
	return &GetAllProductsUseCase{repo: repo}
}

func (uc *GetAllProductsUseCase) Execute() ([]entities.Product, error) {
	return uc.repo.GetAllProducts()
}