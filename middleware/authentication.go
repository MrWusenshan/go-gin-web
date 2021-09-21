package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go-gin-web/common"
	"go-gin-web/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取 Authoritarian header
		tokenString := ctx.GetHeader("Authorization")
		// validate token format
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足!",
			})
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)

		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "无效的用户token",
			})
			ctx.Abort()
			return
		}

		userId := claims.UserId
		db := common.GetDb()
		var user models.User
		db.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "权限不足!",
			})
			ctx.Abort()
			return
		}

		ctx.Set("user", user)
		ctx.Next()
	}
}
