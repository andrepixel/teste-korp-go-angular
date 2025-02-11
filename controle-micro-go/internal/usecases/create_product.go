package usecases

import (
	"controle-micro-go/internal/entities"
	"controle-micro-go/internal/interfaces"
)

type CreateProductUseCase struct {
	repo interfaces.ProductRepositoryInterface
}

func NewCreateProductUseCase(repo interfaces.ProductRepositoryInterface) *CreateProductUseCase {
	return &CreateProductUseCase{repo: repo}
}

func (uc *CreateProductUseCase) Execute(product entities.Product) (entities.Product, error) {
	return uc.repo.CreateProduct(product)
}
