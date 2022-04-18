package Ecode

import "net/http"

var (
	ResponseOkCode   = ErrorCodes{Code: http.StatusOK, Message: "ok"}                    //公共类型响应定义
	ServiceErrorCode = ErrorCodes{Code: http.StatusInternalServerError, Message: "服务异常"} //公共错误响应定义

	ParamErrCode = ErrorCodes{Code: 100, Message: "参数错误"} //公共参数错误响应定义
)
