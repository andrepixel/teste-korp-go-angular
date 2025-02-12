package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/interfaces"
)

type CreateInvoiceUseCase struct {
	repo interfaces.InvoiceRepositoryInterface
}

func NewCreateInvoiceUseCase(repo interfaces.InvoiceRepositoryInterface) *CreateInvoiceUseCase {
	return &CreateInvoiceUseCase{repo: repo}
}

func (uc *CreateInvoiceUseCase) Execute(invoice entities.Invoice) (entities.Invoice, error) {
	return uc.repo.CreateInvoice(invoice)
}
