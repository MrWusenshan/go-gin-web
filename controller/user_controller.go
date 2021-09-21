package controller

import (
	"crypto/bcrypt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go-gin-web/common"
	"go-gin-web/dto"
	"go-gin-web/models"
	"go-gin-web/response"
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
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位!")
		return
	}
	db := common.GetDb()
	if isTelephoneExits(db, telephone) {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码已被注册!")
		return
	}
	if len(password) < 8 || len(password) > 16 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "密码长度必须在8-16位之间!")
		return
	}

	if len(name) == 0 {
		name = utils.RandomUsername(10)
	}
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "用户密码加密出错!")
		return
	}
	newUser := models.User{
		Name:      name,
		Telephone: telephone,
		Password:  string(hashPassword),
	}
	db.Create(&newUser)

	response.SuccessResponse(ctx, nil, "注册成功!")
}

func UserLogin(ctx *gin.Context) {
	telephone := ctx.PostForm("telephone")
	password := ctx.PostForm("password")

	if len(telephone) != 11 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "手机号码必须为11位!")
		return
	}
	db := common.GetDb()
	var user models.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID == 0 {
		response.Response(ctx, http.StatusUnprocessableEntity, 422, nil, "该手机号未注册!")
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		response.FailResponse(ctx, nil, "密码错误!")
		return
	}

	token, err := common.ReleaseToken(user)
	if err != nil {
		response.Response(ctx, http.StatusInternalServerError, 500, nil, "系统出错!")
		log.Printf("token generate err: %v", err)
		return
	}
	response.SuccessResponse(ctx, gin.H{"token": token}, "登录成功!")
}

func UserInfo(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	response.SuccessResponse(ctx, gin.H{"user": dto.ToUserDto(user.(models.User))}, "success!")
}

func isTelephoneExits(db *gorm.DB, telephone string) bool {
	var user models.User
	db.Where("telephone = ? ", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
