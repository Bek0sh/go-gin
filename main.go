package main

import (
	"log"
	"os"
	"project2/pkg/config"
	"project2/pkg/controllers"
	"project2/pkg/repository"
	repoInterfaces "project2/pkg/repository/repoInterfaces"
	"project2/pkg/routes"
	"project2/pkg/service"
	serviceInterfaces "project2/pkg/service/serviceInterfaces"

	"github.com/gin-gonic/gin"
)

var repo repoInterfaces.UserInterface
var services serviceInterfaces.AuthServiceInterface
var controller *controllers.UserController

func init() {
	config.Connect()

	repo = repository.NewRepository(config.DB)
	services = service.NewAuthService(repo)
	controller = controllers.NewUserController(services)

}

func main() {

	log.Print(os.Getenv("ACCESS_TOKEN_PRIVATE_KEY"))

	router := gin.Default()

	r := routes.NewRoutes(*controller, repo)

	r.AuthRoutes(router)

	log.Fatal(router.Run("localhost:8080"))
}
