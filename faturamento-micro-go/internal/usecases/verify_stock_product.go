package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/infrastructure/gateways"
	"fmt"
)

type ValidateProductAvailabilityUseCase struct {
	stock_gateway gateways.StockGateway
}

func NewValidateProductAvailabilityUseCase(stockGateway gateways.StockGateway) *ValidateProductAvailabilityUseCase {
	return &ValidateProductAvailabilityUseCase{
		stock_gateway: stockGateway,
	}
}

func (uc *ValidateProductAvailabilityUseCase) Execute(productsSoldOut []entities.FieldUpdatedProduct) error {
	for _, product := range productsSoldOut {
		product_Reponse, err := uc.stock_gateway.GetProductByID(product.ID)

		if err != nil {
			return fmt.Errorf("erro ao buscar produto: %w", err)
		}

		if product.Available <= 0 {
			return fmt.Errorf("produto %s está sem estoque disponível", product_Reponse.Name)
		}
	}

	return nil
}
