package repositories

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/interfaces"

	"gorm.io/gorm"
)

type InvoiceRepository struct {
	db *gorm.DB
}

func NewInvoiceRepository(db *gorm.DB) interfaces.InvoiceRepositoryInterface {
	return &InvoiceRepository{db: db}
}

func (i *InvoiceRepository) GetAllInvoices() ([]entities.Invoice, error) {
	var invoices []entities.Invoice

	result := i.db.Preload("Items").Find(&invoices) 

	return invoices, result.Error
}

func (i *InvoiceRepository) CreateInvoice(invoice entities.Invoice) (entities.Invoice, error) {
	result := i.db.Create(&invoice)

	return invoice, result.Error
}

func (i *InvoiceRepository) FindInvoiceByID(id string) (entities.Invoice, error) {
	var invoice entities.Invoice

	result := i.db.Preload("Items").First(&invoice, "id = ?", id)

	return invoice, result.Error
}

func (i *InvoiceRepository) UpdateInvoice(id string, invoiceData map[string]interface{}) (entities.Invoice, error) {
	var invoice entities.Invoice

	result := i.db.Model(&invoice).Where("id = ?", id).Updates(invoiceData)

	if result.Error != nil {
		return entities.Invoice{}, result.Error
	}

	return invoice, nil
}

func (i *InvoiceRepository) DeleteInvoice(id string) error {
	result := i.db.Delete(&entities.Invoice{}, "id = ?", id)

	return result.Error
}