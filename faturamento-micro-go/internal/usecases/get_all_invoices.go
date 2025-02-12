package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/interfaces"
)

type GetAllInvoicesUseCase struct {
	repo interfaces.InvoiceRepositoryInterface
}

func NewGetAllInvoicesUseCase(repo interfaces.InvoiceRepositoryInterface) *GetAllInvoicesUseCase {
	return &GetAllInvoicesUseCase{repo: repo}
}

func (uc *GetAllInvoicesUseCase) Execute() ([]entities.Invoice, error) {
	return uc.repo.GetAllInvoices()
}
