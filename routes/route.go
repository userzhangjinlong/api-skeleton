package Route

import (
	"api-skeleton/app/Http/Controller/Api"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodDelete  = "DELETE"
	MethodPatch   = "PATCH"
	MethodConnect = "CONNECT"
	MethodTract   = "TRACE"
	MethodAny     = "ANY"
)

type Route struct {
	Method   string
	Pattern  string
	Callback interface{}
}

// 控制器注入
var (
	IndexController       Api.Index
	LoginController       Api.Login
	UserFriendController  Api.UserFriend
	FriendApplyController Api.FriendApply
	WsServerController    Api.WsServerController
)

//setWebRoute web路由地址
func setWebRoute() map[string][]Route {
	//这里写入所有对应的路由插入
	routes := map[string][]Route{
		//鉴权路由
		"api-login": {
			{MethodGet, "/user/user-friend-list", UserFriendController.UserFriendList},     //好友列表
			{MethodPost, "/user/friend-apply", FriendApplyController.ApplyFriend},          //好友申请
			{MethodPost, "/user/deal-friend-apply", FriendApplyController.DealFriendApply}, //处理好友申请
			{MethodGet, "/user/get-history-msg", UserFriendController.GetHistoryMessage},   //获取聊天历史记录
		},
		//不鉴权路由
		"api-unLogin": {
			{MethodPost, "/login", LoginController.Login},
			{MethodGet, "/index", IndexController.Index},
		},
		//http升级ws路由
		"ws-unLogin": {
			{MethodGet, "/connect", WsServerController.Start},
		},
	}

	return routes
}
