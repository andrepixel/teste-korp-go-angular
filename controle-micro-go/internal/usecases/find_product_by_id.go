package usecases

import (
	"controle-micro-go/internal/entities"
	"controle-micro-go/internal/interfaces"
)

type FindProductByIDUseCase struct {
	repo interfaces.ProductRepositoryInterface
}

func NewFindProductByIDUseCase(repo interfaces.ProductRepositoryInterface) *FindProductByIDUseCase {
	return &FindProductByIDUseCase{repo: repo}
}

func (uc *FindProductByIDUseCase) Execute(id string) (entities.Product, error) {
	return uc.repo.FindProductByID(id)
}
