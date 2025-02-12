package usecases

import (
	"faturamento-micro-go/internal/entities"
	"faturamento-micro-go/internal/infrastructure/gateways"
	"fmt"
)

type UpdateStockUseCase struct {
	stockGateway gateways.StockGateway
}

func NewUpdateStockUseCase(stockGateway gateways.StockGateway) *UpdateStockUseCase {
	return &UpdateStockUseCase{
		stockGateway: stockGateway,
	}
}

func (uc *UpdateStockUseCase) Execute(productsSoldOut []entities.FieldUpdatedProduct) error {
	err := uc.stockGateway.UpdateProduct(productsSoldOut)

	if err != nil {
		return fmt.Errorf("erro ao atualizar estoque: %w", err)
	}

	return nil
}
