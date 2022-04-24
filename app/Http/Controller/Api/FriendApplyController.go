package Api

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Http/Request"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Logic/ApiLogic"
	"api-skeleton/app/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type FriendApply struct {
	Base             *BaseController
	FriendApplyLogic *ApiLogic.FriendApplyLogic
}

//ApplyFriend 发起好友申请
func (fa *FriendApply) ApplyFriend(ctx *gin.Context) {
	userInfo := fa.Base.UserInfo()
	var applyForm ApiRequest.ApplyForm
	valid, errs := Request.BindAndValid(ctx, &applyForm)
	if !valid {
		Util.Error(ctx, Ecode.ParamErrCode.Code, fmt.Sprintf("参数错误：%s", errs))
		return
	}
	err := fa.FriendApplyLogic.ApplyFriend(applyForm, userInfo)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"applyForm": applyForm,
			"err":       err,
		}).Error("好友申请失败")
		Util.Error(ctx, Ecode.FailedCode.Code, "好友申请失败，请重试")
		return
	}
	Util.Success(ctx, "success")
}

//DealFriendApply 处理好友申请
func (fa *FriendApply) DealFriendApply(ctx *gin.Context) {
	dealApplyForm := ApiRequest.DealApplyForm{}
	valid, errs := Request.BindAndValid(ctx, &dealApplyForm)
	if !valid {
		Util.Error(ctx, Ecode.ParamErrCode.Code, fmt.Sprintf("参数错误：%s", errs))
		return
	}

	err := fa.FriendApplyLogic.DealFriendApply(&dealApplyForm)

	if err != nil {
		Util.Error(ctx, Ecode.FailedCode.Code, "服务异常，请重试")
		return
	}

	Util.Success(ctx, "success")

}
