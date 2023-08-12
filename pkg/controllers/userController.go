package controllers

import (
	"net/http"
	"project2/pkg/config"
	"project2/pkg/models"
	iservice "project2/pkg/service/serviceInterfaces"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	service iservice.AuthServiceInterface
}

func NewUserController(service iservice.AuthServiceInterface) *UserController {
	return &UserController{service: service}
}

func (cont *UserController) Register(c *gin.Context) {
	var user models.RegisterUser

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(
			http.StatusNotAcceptable,
			gin.H{
				"status": "failed",
				"error":  err.Error(),
			},
		)

		return
	}

	createdUser, err := cont.service.Register(user)

	if err != nil {
		c.AbortWithStatusJSON(
			http.StatusBadRequest,
			gin.H{
				"status": "failed",
				"error":  err.Error(),
			},
		)
	}
	c.JSON(
		http.StatusOK,
		gin.H{
			"created user": createdUser,
		},
	)
}

func (cont *UserController) SignIn(c *gin.Context) {
	var userInput models.SignInUser

	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(
			http.StatusNonAuthoritativeInfo,
			gin.H{
				"status": "failed",
				"error":  err.Error(),
			},
		)
		return
	}

	config, _ := config.LoadConfig(".")

	accessToken, refreshToken, err := cont.service.SignIn(userInput)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status": "failed",
				"error":  "failed to generate tokens",
			},
		)
		return
	}

	c.SetCookie("access_token", accessToken, config.AccessTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("refresh_token", refreshToken, config.RefreshTokenMaxAge*60, "/", "localhost", false, true)
	c.SetCookie("loged_in", "true", config.AccessTokenMaxAge*60, "/", "localhost", false, false)

	c.JSON(
		http.StatusOK,
		gin.H{
			"status":       "success",
			"access_token": accessToken,
		},
	)

}

func (cont *UserController) LogoutUser(ctx *gin.Context) {
	ctx.SetCookie("access_token", "", -1, "/", "api/v1", false, true)
	ctx.SetCookie("refresh_token", "", -1, "/", "api/v1", false, true)
	ctx.SetCookie("logged_in", "", -1, "/", "api/v1", false, false)

	ctx.JSON(http.StatusOK, gin.H{"status": "success"})

}

func (cont *UserController) Profile(ctx *gin.Context) {
	current_user := ctx.MustGet("current_user")

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":       "success",
			"current_user": current_user,
		},
	)
}
