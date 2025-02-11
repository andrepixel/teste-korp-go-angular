package usecases

import (
	"controle-micro-go/internal/entities"
	"controle-micro-go/internal/interfaces"
)

type UpdateProductUseCase struct {
	repo interfaces.ProductRepositoryInterface
}

func NewUpdateProductUseCase(repo interfaces.ProductRepositoryInterface) *UpdateProductUseCase {
	return &UpdateProductUseCase{repo: repo}
}

func (uc *UpdateProductUseCase) Execute(id string, product map[string]interface{}) (entities.Product, error) {
	return uc.repo.UpdateProduct(id, product)
}
