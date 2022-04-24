package ApiLogic

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Model/Im"
)

type UserFriendLogic struct {
}

type UserFriendJson struct {
	Id       int    `gorm:"column:id" json:"id"`
	Username string `gorm:"column:username" json:"username"`
	Tel      string `gorm:"column:tel" json:"tel"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
}

//GetUserFriendList 获取好友列表
func (ufl *UserFriendLogic) GetUserFriendList(userinfoId int, page int, pageSize int) ([]UserFriendJson, error) {
	var UserFriendJson []UserFriendJson
	var userFriendIds []int64
	where := map[string]interface{}{
		"userId": userinfoId,
		"status": Im.NORMAL_STATUS,
	}
	err := Global.ImDB.Model(Im.UserFriend{}).
		Where(where).
		Pluck("friendUserId", &userFriendIds).Error
	if err != nil {
		return nil, err
	}
	var userModel ApiSkeleton.User
	//获取用户信息
	err = Global.DB.Model(userModel).
		Scopes(userModel.Paginate(page, pageSize)).
		Where("id in ?", userFriendIds).
		Find(&UserFriendJson).Error
	return UserFriendJson, nil
}
