package routes

import (
	"project2/pkg/controllers"
	"project2/pkg/middlware"
	repoInterfaces "project2/pkg/repository/irepository"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controllers.UserController
	controllers.MarketController
	repoInterfaces.UserInterface
}

func NewRoutes(cont controllers.UserController, marketCont controllers.MarketController, repo repoInterfaces.UserInterface) Routes {
	return Routes{cont, marketCont, repo}
}

func (r *Routes) AuthRoutes(router *gin.Engine) {

	router.POST("/auth/register", r.Register)
	router.POST("/auth/sign-in", r.SignIn)

	router.GET("/logout", middlware.CheckAuthMiddleware(r.UserInterface), r.LogoutUser)
	router.GET("/profile", middlware.CheckAuthMiddleware(r.UserInterface), r.Profile)

}

func (r *Routes) MarketRoutes(router *gin.Engine) {
	router.POST("/market/create", middlware.CheckAuthMiddleware(r.UserInterface), r.CreateProduct)
	router.GET("/market/:id", r.GetProductWithId)
	// router.GET("/market/:name", r.GetAllProductsWithName)
	router.GET("/market/", r.GetAllProducts)
	router.DELETE("/market/delete/:id", middlware.CheckAuthMiddleware(r.UserInterface), r.DeleteProductById)
}
