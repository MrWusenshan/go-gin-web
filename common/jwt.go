package common

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"go-gin-web/models"
)

var jwtKey = []byte("my_jwt_key")

type Claims struct {
	UserId uint
	jwt.StandardClaims
}

// ReleaseToken 创建用户token函数
func ReleaseToken(user models.User) (token string, err error) {
	expirationTime := time.Now().Add(7 * 24 * time.Hour)
	claims := &Claims{
		UserId: uint(user.ID),
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(), // token 过期时间
			IssuedAt:  time.Now().Unix(),     // token 发放时间
			Issuer:    "wss",                 // token 签发者
			Subject:   "user token",          // token 主题
		},
	}

	result := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = result.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return token, err
}

// ParseToken 解析token函数
func ParseToken(tokenString string) (*jwt.Token, *Claims, error) {
	claims := &Claims{}

	withClaims, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	return withClaims, claims, err
}
