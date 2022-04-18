package Api

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Http/Request"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Logic/ApiLogic"
	"api-skeleton/app/Util"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Login struct {
}

//Login 登陆注册
func (l *Login) Login(ctx *gin.Context) {
	//参数校验
	var param ApiRequest.LoginForm
	valid, errs := Request.BindAndValid(ctx, &param)
	if !valid {
		Util.Error(ctx, Ecode.ParamErrCode.Code, fmt.Sprintf("参数错误：%s", errs))
		return
	}

	token, err := ApiLogic.LoginLogic(&param)
	if err != nil {
		Util.Error(ctx, Ecode.ParamErrCode.Code, fmt.Sprintf("登陆失败，账号不存在或密码密码错误:%v", err))
		return
	}

	Util.Success(ctx, token)
}
