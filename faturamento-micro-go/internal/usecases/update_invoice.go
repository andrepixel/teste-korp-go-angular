package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/interfaces"
)

type UpdateInvoiceUseCase struct {
	repo interfaces.InvoiceRepositoryInterface
}

func NewUpdateInvoiceUseCase(repo interfaces.InvoiceRepositoryInterface) *UpdateInvoiceUseCase {
	return &UpdateInvoiceUseCase{repo: repo}
}

func (uc *UpdateInvoiceUseCase) Execute(id string, invoice map[string]interface{}) (entities.Invoice, error) {
	return uc.repo.UpdateInvoice(id, invoice)
}
