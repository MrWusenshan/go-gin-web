package utils

import (
	"math/rand"
	"time"
)

var letter = []byte("qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM1234567890")

// RandomUsername 生成随即用户名工具函数
func RandomUsername(length int) (username string) {
	result := make([]byte, length)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letter[rand.Intn(len(letter))]
	}
	return string(result)
}
