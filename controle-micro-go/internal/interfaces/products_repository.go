package interfaces

import "controle-micro-go/internal/entities"

type ProductRepositoryInterface interface {
	GetAllProducts() ([]entities.Product, error)
	CreateProduct(product entities.Product) (entities.Product, error)
	FindProductByID(id string) (entities.Product, error)
	UpdateProduct(id string, product map[string]interface{}) (entities.Product, error)
	DeleteProduct(id string) error
}
