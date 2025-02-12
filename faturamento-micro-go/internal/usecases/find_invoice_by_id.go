package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/interfaces"
)

type FindInvoiceByIDUseCase struct {
	repo interfaces.InvoiceRepositoryInterface
}

func NewFindInvoiceByIDUseCase(repo interfaces.InvoiceRepositoryInterface) *FindInvoiceByIDUseCase {
	return &FindInvoiceByIDUseCase{repo: repo}
}

func (uc *FindInvoiceByIDUseCase) Execute(id string) (entities.Invoice, error) {
	return uc.repo.FindInvoiceByID(id)
}
