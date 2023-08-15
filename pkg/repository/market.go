package repository

import (
	"project2/pkg/models"
	interfaces "project2/pkg/repository/irepository"

	"gorm.io/gorm"
)

type market struct {
	db *gorm.DB
}

func NewMarketRepository(db *gorm.DB) interfaces.MarketInterface {
	return &market{db: db}
}

func (m *market) CreateProduct(product *models.Product) (uint, error) {
	if err := m.db.Create(&product).Error; err != nil {
		return 0, err
	}

	return product.ID, nil
}
func (m *market) GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	if err := m.db.First(&product, "id=?", id).Error; err != nil {
		return &models.Product{}, err
	}

	return &product, nil
}
func (m *market) GetProductByName(name string) ([]models.Product, error) {
	var products []models.Product
	if err := m.db.Where("name=?", name).Find(&products).Error; err != nil {
		return []models.Product{}, err
	}

	return products, nil
}
func (m *market) DeleteProductById(id uint) error {

	if err := m.db.Where("id=?", id).Delete(&models.Product{}).Error; err != nil {
		return err
	}

	return nil
}

func (m *market) GetAllProducts() ([]models.Product, error) {
	var products []models.Product

	if err := m.db.Find(&products).Error; err != nil {
		return products, err
	}

	return products, nil
}
