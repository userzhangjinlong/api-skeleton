package wsRoutes

import "api-skeleton/app/Http/Controller/Im"

var (
	ChatController Im.ChatController
)

type WsRouter struct {
}

func NewWsRouter() *WsRouter {
	return &WsRouter{}
}

func (wr *WsRouter) InitWsRouter() *WsEngine {
	engine := new(WsEngine)
	engine.Group("ws")
	//todo::use 中间件解析token获取用户信息待实现
	engine.AddRoute("to-chat", ChatController.ToChat)
	engine.AddRoute("heart-check", ChatController.HeartCheck)

	return engine
}
