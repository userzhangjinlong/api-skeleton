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
)

//setWebRoute 路由地址
func setWebRoute() map[string][]Route {
	//这里写入所有对应的路由插入
	routes := map[string][]Route{
		//鉴权路由
		"api-login": {
			{MethodGet, "/user/user-friend-list", UserFriendController.UserFriendList},
			{MethodPost, "/user/friend-apply", FriendApplyController.ApplyFriend},
		},
		//不鉴权路由
		"api-unLogin": {
			{MethodPost, "/login", LoginController.Login},
			{MethodGet, "/index", IndexController.Index},
		},
	}

	return routes
}
