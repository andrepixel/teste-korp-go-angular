package repositories

import (
	"controle-micro-go/internal/entities"
	"controle-micro-go/internal/interfaces"
	"errors"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) interfaces.ProductRepositoryInterface {
	return &ProductRepository{db: db}
}

func (r *ProductRepository) GetAllProducts() ([]entities.Product, error) {
	var products []entities.Product

	result := r.db.Find(&products)

	return products, result.Error
}

func (r *ProductRepository) CreateProduct(product entities.Product) (entities.Product, error) {
	result := r.db.Create(&product)

	return product, result.Error
}

func (r *ProductRepository) FindProductByID(id string) (entities.Product, error) {
	var product entities.Product

	result := r.db.First(&product, "id = ?", id)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return entities.Product{}, nil
		}

		return entities.Product{}, result.Error
	}

	return product, nil
}

func (r *ProductRepository) UpdateProduct(id string, productData map[string]interface{}) (entities.Product, error) {
	result := r.db.Model(&entities.Product{}).Where("id = ?", id).Updates(productData)

	if result.Error != nil {
		return entities.Product{}, result.Error
	}

	var updatedProduct entities.Product
	if err := r.db.First(&updatedProduct, "id = ?", id).Error; err != nil {
		return entities.Product{}, err
	}

	return updatedProduct, nil
}

func (r *ProductRepository) DeleteProduct(id string) error {
	result := r.db.Delete(&entities.Product{}, "id = ?", id)

	return result.Error
}
