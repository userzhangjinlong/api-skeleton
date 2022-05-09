package Interface

import (
	"api-skeleton/wsRoutes"
)

type InitRouter interface {
	InitWsRouter() *wsRoutes.WsEngine
}
