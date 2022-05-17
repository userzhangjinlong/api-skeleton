package Api

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Logic/ApiLogic"
	"api-skeleton/app/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserFriend struct {
	Base            *BaseController
	UserFriendLogic *ApiLogic.UserFriendLogic
}

//UserFriendList 好友列表
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

//GetHistoryMessage 获取指定好友聊天记录
func (uf *UserFriend) GetHistoryMessage(ctx *gin.Context) {
	userInfo := uf.Base.UserInfo()
	userinfoId, _ := strconv.Atoi(userInfo.ID)
	toUserId, _ := ctx.GetQuery("toUserId")
	chatUserId, _ := strconv.Atoi(toUserId)
	if toUserId == "" {
		Util.Error(ctx, Ecode.NotFoundCode.Code, Ecode.NotFoundCode.Message)
		return
	}
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", ConstDir.PAGE))
	pageSize, _ := strconv.Atoi(ctx.DefaultQuery("pageSize", ConstDir.PAGE_SIZE))
	historyMsgList, err := uf.UserFriendLogic.GetHistoryMsgList(userinfoId, chatUserId, page, pageSize)
	if err != nil {
		fmt.Println(err)
		Util.Error(ctx, Ecode.FailedCode.Code, Ecode.FailedCode.Message)
		return
	}

	Util.Success(ctx, historyMsgList)
}
