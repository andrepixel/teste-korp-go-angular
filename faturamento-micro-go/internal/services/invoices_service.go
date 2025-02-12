package services

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/infrastructure/gateways"
	"faturamento-micro-go/internal/interfaces"
	"faturamento-micro-go/internal/usecases"
)

type InvoiceService struct {
	getAllInvoicesUseCase  *usecases.GetAllInvoicesUseCase
	createInvoiceUseCase   *usecases.CreateInvoiceUseCase
	findInvoiceByIDUseCase *usecases.FindInvoiceByIDUseCase
	updateInvoiceUseCase   *usecases.UpdateInvoiceUseCase
	deleteInvoiceUseCase   *usecases.DeleteInvoiceUseCase
}

func NewInvoiceService(repo interfaces.InvoiceRepositoryInterface) *InvoiceService {
	return &InvoiceService{
		getAllInvoicesUseCase:  usecases.NewGetAllInvoicesUseCase(repo),
		createInvoiceUseCase:   usecases.NewCreateInvoiceUseCase(repo),
		findInvoiceByIDUseCase: usecases.NewFindInvoiceByIDUseCase(repo),
		updateInvoiceUseCase:   usecases.NewUpdateInvoiceUseCase(repo),
		deleteInvoiceUseCase:   usecases.NewDeleteInvoiceUseCase(repo),
	}
}

func (s *InvoiceService) GetAllInvoices() ([]entities.Invoice, error) {
	return s.getAllInvoicesUseCase.Execute()
}

func (s *InvoiceService) CreateInvoice(invoice entities.Invoice) (entities.Invoice, error) {
	products := []entities.FieldUpdatedProduct{}

	_, err := s.FindInvoiceByID(invoice.ID.String())

	if err != nil {
		return entities.Invoice{}, err
	}

	for _, item := range invoice.Items {
		var prod entities.FieldUpdatedProduct

		prod.ID = item.ID.String()

		prod.Available = item.Availabe

		products = append(products, prod)
	}

	stock_gateway, _ := gateways.NewStockGateway()

	err2 := usecases.NewValidateProductAvailabilityUseCase(*stock_gateway).Execute(products)

	if err2 != nil {
		return entities.Invoice{}, err2
	}

	err3 := usecases.NewUpdateStockUseCase(*stock_gateway).Execute(products)

	if err3 != nil {
		return entities.Invoice{}, err3
	}

	entity, err4 := s.createInvoiceUseCase.Execute(invoice)

	if err4 != nil {
		return entities.Invoice{}, err4
	}

	return entity, nil
}

func (s *InvoiceService) FindInvoiceByID(id string) (entities.Invoice, error) {
	return s.findInvoiceByIDUseCase.Execute(id)
}

func (s *InvoiceService) DeleteInvoice(idInvoice string) (entities.Invoice, error) {
	products := []entities.FieldUpdatedProduct{}

	invoice, err := s.FindInvoiceByID(idInvoice)

	if err != nil {
		return entities.Invoice{}, err
	}

	for _, item := range invoice.Items {
		var prod entities.FieldUpdatedProduct

		prod.ID = item.ID.String()

		prod.Available = item.Availabe

		products = append(products, prod)
	}

	stock_gateway, _ := gateways.NewStockGateway()

	product_response, err2 := usecases.NewGetAllProductsInStockUseCase(*stock_gateway).Execute(products)

	if err2 != nil {
		return entities.Invoice{}, err2
	}

	var products_with_available_changed = products

	for i := range products_with_available_changed {
		for _, product_response := range product_response {
			for _, product := range products {
				products_with_available_changed[i].Available = product_response.Available + product.Available
			}
		}
	}

	err3 := usecases.NewUpdateStockUseCase(*stock_gateway).Execute(products_with_available_changed)

	if err3 != nil {
		return entities.Invoice{}, err3
	}

	entity, err4 := s.createInvoiceUseCase.Execute(invoice)

	if err4 != nil {
		return entities.Invoice{}, err4
	}

	return entity, nil
}