package Api

import (
	"api-skeleton/app/Util"
	"github.com/gin-gonic/gin"
)

type UserFriend struct {
	//继承父类Base 做其他逻辑处理
	Base *BaseController
}

func (uf *UserFriend) UserFriendList(ctx *gin.Context) {
	Util.Success(ctx, []string{"灰常滴成功"})
}
