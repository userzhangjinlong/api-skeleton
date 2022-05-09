package interfaceDepend

import (
	"api-skeleton/app/Interface"
	"api-skeleton/wsRoutes"
)

func NewWsRouterDepend() *wsRoutes.WsRouter {
	return wsRoutes.NewWsRouter()
}

func InitWsRouter(router Interface.InitRouter) *wsRoutes.WsEngine {
	return router.InitWsRouter()
}
