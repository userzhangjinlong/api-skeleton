package ApiLogic

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Model/Im"
	"sort"
)

type UserFriendLogic struct {
}

type UserFriendJson struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Tel      string `gorm:"column:tel" json:"tel"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
}

type HistoryMsgJson struct {
	FromUserId int    `gorm:"column:fromUserId" json:"fromUserId"`
	ToUserId   string `gorm:"column:toUserId" json:"toUserId"`
	Content    string `gorm:"column:content" json:"content"`
	MsgImg     string `gorm:"column:msgImg" json:"msgImg"`
	SendTime   string `gorm:"column:sendTime" json:"sendTime"`
}

type HistoryMsgResponse struct {
	MsgList    []HistoryMsgJson `json:"msgList"`
	ToUserInfo UserFriendJson   `json:"toUserInfo"`
}

//GetUserFriendList 获取好友列表
func (ufl *UserFriendLogic) GetUserFriendList(userinfoId int, page int, pageSize int) ([]UserFriendJson, error) {
	var UserFriendJson []UserFriendJson
	var userFriendIds []int64
	tmpFriendIds := make([]struct {
		UserId       int64 `gorm:"column:userId"`
		FriendUserId int64 `gorm:"column:friendUserId"`
	}, 0)
	err := Global.ImDB.Model(Im.UserFriend{}).
		Where("(userId = ? or friendUserId = ?) and status = ?", userinfoId, userinfoId, Im.NORMAL_STATUS).
		Select("userId, friendUserId").
		Find(&tmpFriendIds).Error
	if err != nil {
		return nil, err
	}
	for _, friends := range tmpFriendIds {
		if friends.FriendUserId != int64(userinfoId) {
			userFriendIds = append(userFriendIds, friends.FriendUserId)
		} else {
			userFriendIds = append(userFriendIds, friends.UserId)
		}
	}
	var userModel ApiSkeleton.User
	//获取用户信息
	err = Global.DB.Model(userModel).
		Scopes(userModel.Paginate(page, pageSize)).
		Where("id in ?", userFriendIds).
		Find(&UserFriendJson).Error
	if err != nil {
		return nil, err
	}

	for k, v := range UserFriendJson {
		if v.Avatar == "" {
			UserFriendJson[k].Avatar = "https://cdn.learnku.com//uploads/communities/sNljssWWQoW6J88O9G37.png!/both/44x44"
		}
	}
	return UserFriendJson, nil
}

//GetHistoryMsgList 消息记录列表
func (ufl *UserFriendLogic) GetHistoryMsgList(fromUserId, toUserId, page, pageSize int) (HistoryMsgResponse, error) {

	HistoryMsgList := HistoryMsgResponse{
		MsgList:    []HistoryMsgJson{},
		ToUserInfo: UserFriendJson{},
	}

	var imMsgMode Im.ImMsg
	var userModel ApiSkeleton.User

	err := Global.ImDB.Model(imMsgMode).
		Scopes(imMsgMode.Paginate(page, pageSize)).
		Where("fromUserId = ? or toUserId = ?", fromUserId, fromUserId).
		Order("id DESC").
		Find(&HistoryMsgList.MsgList).Error

	if err != nil {
		return HistoryMsgList, err
	}

	//查询到的聊天数据顺序反转
	sort.Slice(HistoryMsgList.MsgList, func(i, j int) bool {
		return HistoryMsgList.MsgList[i].SendTime < HistoryMsgList.MsgList[j].SendTime
	})
	//获取指定用户信息
	whereUser := map[string]interface{}{
		"id": toUserId,
	}
	err = Global.DB.Model(userModel).
		Where(whereUser).
		First(&HistoryMsgList.ToUserInfo).Error
	if err != nil {
		return HistoryMsgList, err
	}

	return HistoryMsgList, nil
}
