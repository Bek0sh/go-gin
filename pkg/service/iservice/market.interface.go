package interfaces

import "project2/pkg/models"

type MarketServiceInterface interface {
	CreateProduct(*models.ProductInput) (uint, error)
	GetProductById(id uint) (*models.Product, error)
	GetProductByName(name string) ([]models.Product, error)
	DeleteProductById(id uint) error
	GetAllProducts() ([]models.Product, error)
}
