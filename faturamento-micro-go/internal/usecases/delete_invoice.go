package usecases

import (
	"faturamento-micro-go/internal/interfaces"
)

type DeleteInvoiceUseCase struct {
	repo interfaces.InvoiceRepositoryInterface
}

func NewDeleteInvoiceUseCase(repo interfaces.InvoiceRepositoryInterface) *DeleteInvoiceUseCase {
	return &DeleteInvoiceUseCase{repo: repo}
}

func (uc *DeleteInvoiceUseCase) Execute(id string) error {
	return uc.repo.DeleteInvoice(id)
}
