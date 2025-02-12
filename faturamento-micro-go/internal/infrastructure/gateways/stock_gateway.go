package gateways

import (
	"encoding/json"
	"faturamento-micro-go/internal/entities"
	"fmt"
	"os"

	"github.com/go-resty/resty/v2"
)

type StockGateway struct {
	client *resty.Client
	apiURL string
}

func NewStockGateway() (*StockGateway, error) {
	apiURL := os.Getenv("ESTOQUE_SERVICE_URL")

	if apiURL == "" {
		return nil, fmt.Errorf("variável de ambiente ESTOQUE_SERVICE_URL não definida")
	}

	return &StockGateway{
		client: resty.New(),
		apiURL: apiURL,
	}, nil
}
func (g *StockGateway) GetProductByID(productID string) (*entities.ProductResponse, error) {
	url := fmt.Sprintf("%s/?id=%s", g.apiURL, productID)

	resp, err := g.client.R().
		SetHeader("Content-Type", "application/json").
		Get(url)

	if err != nil {
		return nil, fmt.Errorf("erro ao fazer requisição: %w", err)
	}

	if resp.StatusCode() != 200 {
		return nil, fmt.Errorf("erro na resposta do servidor: %d", resp.StatusCode())
	}

	var product entities.ProductResponse

	if err := json.Unmarshal(resp.Body(), &product); err != nil {
		return nil, fmt.Errorf("erro ao decodificar JSON: %w", err)
	}

	return &product, nil
}

func (g *StockGateway) UpdateProduct(productsSoldOut []entities.FieldUpdatedProduct) error {
	for _, product := range productsSoldOut {
		url := fmt.Sprintf("%s/%s", g.apiURL, product.ID)

		resp, err := g.client.R().
			SetHeader("Content-Type", "application/json").
			SetBody(product.Available).
			Patch(url)

		if err != nil {
			return fmt.Errorf("erro ao fazer requisição: %w", err)
		}

		if resp.StatusCode() != 200 {
			return fmt.Errorf("erro na resposta do servidor: %d", resp.StatusCode())
		}
	}

	return nil
}
