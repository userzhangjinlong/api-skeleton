package Api

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request"
	"api-skeleton/app/Model/Im"
	"api-skeleton/app/Util"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

type FriendApply struct {
	Base *BaseController
}

type ApplyForm struct {
	FriendUid int64  `form:"friendUid" binding:"required,min=0"`
	Desc      string `form:"desc" binding:"required,max=200"`
}

var (
	FriendApplyModel *Im.FriendApply
)

//ApplyFriend 发起好友申请
func (fa *FriendApply) ApplyFriend(ctx *gin.Context) {
	userInfo := fa.Base.UserInfo(ctx)
	var applyForm ApplyForm
	valid, errs := Request.BindAndValid(ctx, &applyForm)
	if !valid {
		Util.Error(ctx, Ecode.ParamErrCode.Code, fmt.Sprintf("参数错误：%s", errs))
		return
	}

	err := Global.DB.Clauses(dbresolver.Use(ConstDir.IM)).
		Where("userId = ? and friendUserId = ?", userInfo.ID, &applyForm.FriendUid).
		First(&FriendApplyModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		Util.Error(ctx, Ecode.ServiceErrorCode.Code, fmt.Sprintf("数据查询异常：%s", err))
		return
	}

	var returnVal interface{}
	if FriendApplyModel.Id == 0 {
		returnVal = []string{}
	} else {
		returnVal = FriendApplyModel
	}

	Util.Success(ctx, returnVal)
}
