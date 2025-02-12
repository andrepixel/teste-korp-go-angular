package interfaces

import "faturamento-micro-go/internal/entities"

type InvoiceRepositoryInterface interface {
	GetAllInvoices() ([]entities.Invoice, error)
	CreateInvoice(invoice entities.Invoice) (entities.Invoice, error)
	FindInvoiceByID(id string) (entities.Invoice, error)
	UpdateInvoice(id string, invoice map[string]interface{}) (entities.Invoice, error)
	DeleteInvoice(id string) error
}
