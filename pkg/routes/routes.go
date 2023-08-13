package routes

import (
	"project2/pkg/controllers"
	"project2/pkg/middlware"
	repoInterfaces "project2/pkg/repository/irepository"

	"github.com/gin-gonic/gin"
)

type Routes struct {
	controllers.UserController
	repoInterfaces.UserInterface
}

func NewRoutes(cont controllers.UserController, repo repoInterfaces.UserInterface) Routes {
	return Routes{cont, repo}
}

func (r *Routes) AuthRoutes(router *gin.Engine) {

	router.POST("/auth/register", r.Register)
	router.POST("/auth/sign-in", r.SignIn)

	router.GET("/logout", middlware.CheckAuthMiddleware(r.UserInterface), r.LogoutUser)
	router.GET("/profile", middlware.CheckAuthMiddleware(r.UserInterface), r.Profile)

}
