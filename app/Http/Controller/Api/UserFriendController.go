package Api

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Logic/ApiLogic"
	"api-skeleton/app/Util"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserFriend struct {
	Base            *BaseController
	UserFriendLogic *ApiLogic.UserFriendLogic
}

func (uf *UserFriend) UserFriendList(ctx *gin.Context) {
	userInfo := uf.Base.UserInfo()
	userinfoId, _ := strconv.Atoi(userInfo.ID)
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", ConstDir.PAGE))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", ConstDir.PAGE_SIZE))
	userFriendList, err := uf.UserFriendLogic.GetUserFriendList(userinfoId, page, pageSize)
	if err != nil {
		Util.Error(ctx, Ecode.FailedCode.Code, Ecode.FailedCode.Message)
		return
	}
	Util.Success(ctx, userFriendList)
}
