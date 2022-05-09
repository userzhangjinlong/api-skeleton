package wsRoutes

import (
	"api-skeleton/grateway"
)

type WsEngine struct {
	WsGroup
}

func NewWsEngine() *WsEngine {
	engine := &WsEngine{
		WsGroup{
			basePath:           "",
			MiddlewareHandlers: nil,
			RouteHandlers:      nil,
		},
	}
	engine.WsGroup.engine = engine
	return engine
}

func (we *WsEngine) Run(path string, req *grateway.WsMsgReq) {
	//对应
	we.WsGroup.Exec(path, req)
}
