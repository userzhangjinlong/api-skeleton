package Api

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Util"
	"github.com/gin-gonic/gin"
)

type BaseController struct {
}

//UserInfo 用户信息获取
func (b *BaseController) UserInfo(ctx *gin.Context) *Util.UserClaims {
	userInfo, err := ctx.Get(ConstDir.AUTH_USER)
	if !err {
		//todo::这里为终止响应问题待处理
		Util.Error(ctx, Ecode.NotFoundCode.Code, "用户信息获取失败")
		return nil
	}

	return userInfo.(*Util.UserClaims)
}
