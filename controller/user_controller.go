package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-web/common"
	"go-gin-web/models"
	"go-gin-web/utils"
	"gorm.io/gorm"
)

func UserRegister(ctx *gin.Context) {
	// 接收参数
	name := ctx.PostForm("name")
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	// 校验参数
	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号码必须为11位!",
		})
		return
	}
	db := common.GetDb()
	if isTelephoneExits(db, telephone) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号码已被注册!",
		})
		return
	}
	if len(password) < 8 || len(password) > 16 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "密码长度必须在8-16位之间!",
		})
		return
	}

	if len(name) == 0 {
		name = utils.RandomUsername(10)
	}

	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  password,
	}
	db.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功!",
	})
}

func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ? ", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
