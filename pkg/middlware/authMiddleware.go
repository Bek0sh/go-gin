package middlware

import (
	"net/http"
	"project2/pkg/config"
	interfaces "project2/pkg/repository/repoInterfaces"
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

		id, err := utils.VerifyToken(access_token, config.AccessTokenPubliceKey)

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status": "failed",
					"error":  err.Error(),
				},
			)
		}

		user, err := repo.FindUserById(id.(int))

		if err != nil {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				gin.H{
					"status":  "fail",
					"message": "The user belonging to this token no logger exists",
				},
			)
			return
		}

		ctx.Set("current_user", user)
		ctx.Next()
	}
}
