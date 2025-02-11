package services

import (
	"controle-micro-go/internal/entities"
	"controle-micro-go/internal/interfaces"
	"controle-micro-go/internal/usecases"
	"encoding/json"
	"fmt"
)

type ProductService struct {
	getAllProductsUseCase  *usecases.GetAllProductsUseCase
	createProductUseCase   *usecases.CreateProductUseCase
	findProductByIDUseCase *usecases.FindProductByIDUseCase
	updateProductUseCase   *usecases.UpdateProductUseCase
	deleteProductUseCase   *usecases.DeleteProductUseCase
}

func NewProductService(repo interfaces.ProductRepositoryInterface) *ProductService {
	return &ProductService{
		getAllProductsUseCase:  usecases.NewGetAllProductsUseCase(repo),
		createProductUseCase:   usecases.NewCreateProductUseCase(repo),
		findProductByIDUseCase: usecases.NewFindProductByIDUseCase(repo),
		updateProductUseCase:   usecases.NewUpdateProductUseCase(repo),
		deleteProductUseCase:   usecases.NewDeleteProductUseCase(repo),
	}
}

func (s *ProductService) GetAllProducts() ([]entities.Product, error) {
	return s.getAllProductsUseCase.Execute()
}

func (s *ProductService) CreateProduct(product entities.Product) (entities.Product, error) {
	_, err := s.FindProductByID(product.ID.String())

	if err != nil {
		return entities.Product{}, err
	}

	entity, err2 := s.createProductUseCase.Execute(product)

	if err2 != nil {
		return entities.Product{}, err2
	}

	return entity, nil
}

func (s *ProductService) FindProductByID(id string) (entities.Product, error) {
	return s.findProductByIDUseCase.Execute(id)
}

func (s *ProductService) UpdateProduct(id string, product entities.Product) (entities.Product, error) {
	productSearch, err := s.FindProductByID(id)

	if err != nil {
		return entities.Product{}, err
	}

	var result map[string]interface{}

	jsonObject, _ := json.Marshal(product)

	json.Unmarshal(jsonObject, &result)

	delete(result, "id")
	delete(result, "created_at")

	for key, value := range result {
		switch v := value.(type) {
		case string:
			if v == "" || v == "00000000-0000-0000-0000-000000000000" {
				fmt.Println("Removendo:", key, "=", v)
				delete(result, key)
			}

			if key == "CreatedAt" || key == "UpdatedAt" {
				fmt.Println("Removendo timestamp:", key, "=", v)
				delete(result, key)
			}
		case float64:
			if v == 0 {
				fmt.Println("Removendo:", key, "=", v)
				delete(result, key)
			}
		case nil:
			fmt.Println("Removendo:", key, "= nil")
			delete(result, key)
		}
	}

	return s.updateProductUseCase.Execute(productSearch.ID.String(), result)
}

func (s *ProductService) DeleteProduct(id string) (entities.Product, error) {
	entity, err := s.FindProductByID(id)

	if err != nil {
		return entities.Product{}, err
	}

	err2 := s.deleteProductUseCase.Execute(id)

	if err2 != nil {
		return entities.Product{}, err
	}

	return entity, nil
}
