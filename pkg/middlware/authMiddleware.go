package middlware

import (
	"net/http"
	"project2/pkg/config"
	"project2/pkg/models"
	interfaces "project2/pkg/repository/irepository"
	"project2/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func CheckAuthMiddleware(repo interfaces.UserInterface) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var access_token string

		cookie, err := ctx.Cookie("access_token")
		authorization := ctx.GetHeader("Authorization")

		fields := strings.Fields(authorization)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{"status": "fail", "message": "You are not logged in"},
			)
			return
		}

		config, _ := config.LoadConfig(".")

		id, err := utils.VerifyToken(access_token, config.AccessTokenPublicKey)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status": "failed",
					"error":  err.Error(),
				},
			)
		}

		user, err := repo.FindUserById(uint(id.(float64)))

		current_user := &models.ResponseUser{
			Model:   user.Model,
			Name:    user.Name,
			Surname: user.Surname,
			Email:   user.Email,
		}

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  "fail",
					"message": "The user belonging to this token no longer exists",
				},
			)
			return
		}

		ctx.Set("current_user", current_user)
		ctx.Next()
	}
}
