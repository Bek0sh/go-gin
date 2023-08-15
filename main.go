package main

import (
	"log"
	"project2/pkg/config"
	"project2/pkg/controllers"
	"project2/pkg/repository"
	repoInterfaces "project2/pkg/repository/irepository"
	"project2/pkg/routes"
	"project2/pkg/service"
	serviceInterfaces "project2/pkg/service/iservice"

	"github.com/gin-gonic/gin"
)

var repo repoInterfaces.UserInterface
var marketRepo repoInterfaces.MarketInterface

var services serviceInterfaces.AuthServiceInterface
var marketService serviceInterfaces.MarketServiceInterface

var controller *controllers.UserController
var marketController *controllers.MarketController

func init() {
	config.Connect()

	repo = repository.NewRepository(config.DB)
	marketRepo = repository.NewMarketRepository(config.DB)

	services = service.NewAuthService(repo)
	marketService = service.NewMarketService(marketRepo)

	controller = controllers.NewUserController(services)
	marketController = controllers.NewMarketController(marketService)

}

func main() {

	router := gin.Default()

	r := routes.NewRoutes(*controller, *marketController, repo)

	r.AuthRoutes(router)
	r.MarketRoutes(router)

	log.Fatal(router.Run("localhost:8080"))
}
