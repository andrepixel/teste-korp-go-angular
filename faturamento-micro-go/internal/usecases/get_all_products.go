package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/infrastructure/gateways"
	"fmt"
)

type GetAllProductsInStock struct {
	stock_gateway gateways.StockGateway
}

func NewGetAllProductsInStockUseCase(stockGateway gateways.StockGateway) *GetAllProductsInStock {
	return &GetAllProductsInStock{
		stock_gateway: stockGateway,
	}
}

func (uc *GetAllProductsInStock) Execute(productsSoldOut []entities.FieldUpdatedProduct) ([]entities.FieldUpdatedProduct, error) {
	var productsRetorned = productsSoldOut

	for _, product := range productsSoldOut {
		product_Reponse, err := uc.stock_gateway.GetProductByID(product.ID)

		if err != nil {
			return []entities.FieldUpdatedProduct{}, fmt.Errorf("erro ao buscar produto: %w", err)
		}

		for i := range productsRetorned {
			if productsRetorned[i].ID == product.ID {
				productsRetorned[i].Available = product_Reponse.Available
			}
		}
	}

	return productsRetorned, nil
}
