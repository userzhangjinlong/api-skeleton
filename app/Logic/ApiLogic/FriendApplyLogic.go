package ApiLogic

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Http/Request/ApiRequest"
	"api-skeleton/app/Model/Im"
	"api-skeleton/app/Util"
	"errors"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
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
	err := Global.ImDB.
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
	err = Global.ImDB.Create(&FriendApplyModel).Error
	if err != nil {
		return err
	}
	return nil
}

//DealFriendApply 处理申请请求
func (fal *FriendApplyLogic) DealFriendApply(form *ApiRequest.DealApplyForm) error {

	//获取申请的信息
	applyInfo := Im.FriendApply{}
	Global.ImDB.Where("id = ?", form.ApplyId).First(&applyInfo)
	if applyInfo.Id == 0 {
		return errors.New("申请信息不存在请重试")
	}
	transaction := Global.ImDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			logrus.WithFields(
				logrus.Fields{
					"applyForm": form,
					"err":       r,
				}).Error("捕获到异常")
			transaction.Rollback()
		}
	}()

	//处理好友申请
	err := transaction.
		Model(&Im.FriendApply{}).
		Where("id = ?", form.ApplyId).
		Update("status", form.Status).Error
	if err != nil {
		transaction.Rollback()
		logrus.WithFields(logrus.Fields{
			"applyForm": form,
			"err":       err,
		}).Error("处理请求失败")
		return err
	}

	if form.Status == Im.PASS_STATUS {
		//添加好友关系
		userFriendData := Im.UserFriend{}
		userFriendData.UserId = applyInfo.UserId
		userFriendData.FriendUserId = applyInfo.FriendUserId
		userFriendData.Status = Im.NORMAL_STATUS
		err = transaction.
			Create(&userFriendData).Error
		if err != nil {
			transaction.Rollback()
			logrus.WithFields(logrus.Fields{
				"userFriendData": userFriendData,
				"err":            err,
			}).Error("添加好友关系异常")
			return err
		}
	}

	err = transaction.Commit().Error
	if err != nil {
		transaction.Rollback()
		return err
	}

	return nil
}
