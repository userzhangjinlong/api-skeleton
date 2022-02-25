package Ecode

import "net/http"

var (
	//公共类型响应定义
	ResponseOkCode = &ErrorCodes{Code: http.StatusOK, Message: "ok"}
)
