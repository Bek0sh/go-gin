package interfaces

import "project2/pkg/models"

type MarketInterface interface {
	CreateProduct(product *models.ProductInput) (uint, error)
	GetProductById(id uint) (*models.Product, error)
	GetProductByName(name string) ([]models.Product, error)
	GetAllProducts() ([]models.Product, error)
	DeleteProductById(id uint) error
}
