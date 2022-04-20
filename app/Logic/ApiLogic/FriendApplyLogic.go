package ApiLogic

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Model/Im"
	"api-skeleton/app/Util"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
	"strconv"
	"time"
)

type FriendApplyLogic struct {
}

var (
	FriendApplyModel *Im.FriendApply
)

func (fal *FriendApplyLogic) ApplyFriend(applyForm ApiRequest.ApplyForm, userInfo *Util.UserClaims) error {
	err := Global.DB.Clauses(dbresolver.Use(ConstDir.IM)).
		Where("userId = ? and friendUserId = ?", userInfo.ID, &applyForm.FriendUid).
		First(&FriendApplyModel).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	//创建好友申请
	FriendApplyModel.FriendUserId = applyForm.FriendUid
	FriendApplyModel.UserId, _ = strconv.ParseInt(userInfo.ID, 10, 64)
	//申请默认三天时间自动过期
	FriendApplyModel.ExpireTime = time.Now().AddDate(0, 0, 3).Unix()
	FriendApplyModel.Desc = applyForm.Desc
	err = Global.DB.Clauses(dbresolver.Use(ConstDir.IM)).Create(&FriendApplyModel).Error
	if err != nil {
		return err
	}
	return nil
}
