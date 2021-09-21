package controller

import (
	"crypto/bcrypt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-web/common"
	"go-gin-web/dto"
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
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "用户密码加密出错了!",
		})
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	db.Create(&newUser)

	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "注册成功!",
	})
}

func UserLogin(ctx *gin.Context) {
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "手机号码必须为11位!",
		})
		return
	}
	db := common.GetDb()
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{
			"code":    422,
			"message": "该手机号未注册!",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "密码错误!",
		})
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "系统出错!",
		})
		log.Printf("token generate err: %v", err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{
			"token": token,
		},
		"message": "登录成功!",
	})
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{
		"code":    200,
		"data":    gin.H{"user": dto.ToUserDto(user.(models.User))},
		"message": "success!",
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
