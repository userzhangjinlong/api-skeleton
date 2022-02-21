package Route

import (
	"api-skeleton/app/Http/Controller/Api"
	"api-skeleton/app/Http/Middleware"
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
	Method     string
	Pattern    string
	Callback   interface{}
	Middleware interface{}
}

// 控制器注入
var (
	IndexController Api.Index
	LoginController Api.Login
)

func setWebRoute() map[string][]Route {
	//这里写入所有对应的路由插入
	routes := map[string][]Route{
		"v1": {
			{MethodGet, "/index", IndexController.Index, Middleware.Auth()},
		},
		"api": {
			{MethodPost, "/login", LoginController.Login, nil},
		},
	}

	return routes
}
