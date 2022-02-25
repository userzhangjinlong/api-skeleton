package Util

import (
	"api-skeleton/app/Ecode"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

//Success 成功抛出正常信息
func Success(ctx *gin.Context, data interface{}) {
	response := Response{
		Code:    Ecode.ResponseOkCode.Code,
		Message: Ecode.ResponseOkCode.Message,
		Data:    data,
	}
	ctx.JSON(http.StatusOK, response)
}

func Error(ctx *gin.Context, errCode int, msg string) {
	response := Response{Code: errCode, Message: msg}
	ctx.JSON(http.StatusOK, response)
	ctx.Abort()
	return
}
