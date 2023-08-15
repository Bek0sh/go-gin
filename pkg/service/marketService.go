package service

import (
	"errors"
	"project2/pkg/models"
	irepository "project2/pkg/repository/irepository"
	iservice "project2/pkg/service/iservice"
)

type marketService struct {
	repo irepository.MarketInterface
}

func NewMarketService(repo irepository.MarketInterface) iservice.MarketServiceInterface {
	return &marketService{repo: repo}
}

func (service *marketService) CreateProduct(product *models.ProductInput) (uint, error) {

	user := models.User{
		Model: currentUser.Model,
	}

	createdProd := &models.Product{
		Name:    product.Name,
		Price:   product.Price,
		User_ID: user.ID,
	}

	id, err := service.repo.CreateProduct(createdProd)

	if err != nil {
		return 0, err
	}

	return id, nil
}

func (service *marketService) GetProductById(id uint) (*models.Product, error) {
	return service.repo.GetProductById(id)
}

func (service *marketService) GetAllProducts() ([]models.Product, error) {
	return service.repo.GetAllProducts()
}

func (service *marketService) GetProductByName(name string) ([]models.Product, error) {
	return service.repo.GetProductByName(name)
}

func (service *marketService) DeleteProductById(id uint) error {
	product, err := service.GetProductById(id)

	if err != nil {
		return err
	}

	if product.User_ID != currentUser.ID {
		return errors.New("you can not delete others product")
	}

	if err := service.repo.DeleteProductById(id); err != nil {
		return err
	}

	return nil

}
