package ApiRequest

type ApplyForm struct {
	FriendUid int64  `form:"friendUid" binding:"required,min=0"`
	Desc      string `form:"desc" binding:"required,max=200"`
}
