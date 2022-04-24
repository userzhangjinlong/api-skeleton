package Api

import (
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Util"
)

type BaseController struct {
}

//UserInfo 用户信息获取
func (b *BaseController) UserInfo() *Util.UserClaims {

	return ApiSkeleton.UserInfo
}
