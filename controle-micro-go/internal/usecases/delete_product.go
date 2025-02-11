package usecases

import (
	"controle-micro-go/internal/interfaces"
)

type DeleteProductUseCase struct {
	repo interfaces.ProductRepositoryInterface
}

func NewDeleteProductUseCase(repo interfaces.ProductRepositoryInterface) *DeleteProductUseCase {
	return &DeleteProductUseCase{repo: repo}
}

func (uc *DeleteProductUseCase) Execute(id string) error {
	return uc.repo.DeleteProduct(id)
}
