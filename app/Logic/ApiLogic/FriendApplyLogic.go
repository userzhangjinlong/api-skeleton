package ApiLogic

import (
	"api-skeleton/app/ConstDir"
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Model/Im"
	"api-skeleton/app/Util"
	"github.com/sirupsen/logrus"
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

//ApplyFriend 好友申请
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

//DealFriendApply 处理申请请求
func (fal *FriendApplyLogic) DealFriendApply(form *ApiRequest.DealApplyForm) error {
	err := Global.DB.Clauses(dbresolver.Use(ConstDir.IM)).
		Model(&Im.FriendApply{}).
		Where("id = ?", form.ApplyId).
		Update("status", form.Status).Error

	if err != nil {
		logrus.WithFields(logrus.Fields{
			"applyForm": form,
			"err":       err,
		}).Error("处理请求失败")
		return err
	}

	//同意好友申请向好友关系表写入好友关系记录 (事务)

	return nil
}
