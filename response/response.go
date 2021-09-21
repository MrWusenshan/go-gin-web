package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Response 普通返回
func Response(ctx *gin.Context, httpStatusCode, code int, data gin.H, message string) {
	ctx.JSON(httpStatusCode, gin.H{
		"code":    code,
		"data":    data,
		"message": message,
	})
}

// SuccessResponse 成功结果返回
func SuccessResponse(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, 200, data, message)
}

// FailResponse 失败结果返回
func FailResponse(ctx *gin.Context, data gin.H, message string) {
	Response(ctx, http.StatusOK, 400, data, message)
}
