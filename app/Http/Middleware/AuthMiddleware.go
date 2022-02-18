package Middleware

import (
	"api-skeleton/app/Util"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := getToken(ctx)
		if token == "" {
			Util.Error(ctx, 401, "请登录！")
			return
		}

		//token鉴权解析是否成功判断是否登陆成功或过期
		userClaims, err := Util.ParseToken(token)
		if userClaims == nil || err != nil {
			Util.Error(ctx, 401, "token解析异常或者登陆失效")
			return
		} else {
			//设置保存用户信息
			ctx.Set("User", userClaims)
		}

		ctx.Next()
	}
}

//getToken 从请求中获取token
func getToken(ctx *gin.Context) string {
	token := ctx.Request.Header.Get("token")
	if token == "" {
		//获取请求参数里面是否携带token
		param, _ := Util.GetFormParam(ctx)
		token = param["token"]
	}

	return token
}
