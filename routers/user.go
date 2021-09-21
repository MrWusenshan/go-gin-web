package routers

import (
	"github.com/gin-gonic/gin"
	"go-gin-web/controller"
)

// User 用户功能模块相关路由
func User(router *gin.Engine) {
	router.POST("/api/auth/register", controller.UserRegister)
	router.POST("/api/auth/login", controller.UserLogin)
}
