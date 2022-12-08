package middleware

import (
	"gblog-server/common"
	"gblog-server/model"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.Request.Header.Get("Authorization")

		if tokenString == "" {
			ctx.JSON(http.StatusOK, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		if tokenString == "" || len(tokenString) < 7 || !strings.HasPrefix(tokenString, "Bearer") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code": 401,
				"msg":  "权限不足",
			})

			ctx.Abort()
			return
		}

		userId := claims.UserId
		DB := common.GetDB()
		var user model.User

		DB.Where("id = ?", userId).First(&user)
		ctx.Set("user", user)
		ctx.Next()
	}
}
