package Api

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Util"
	"api-skeleton/app/interfaceDepend"
	ws2 "api-skeleton/bootstrap/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsServerController struct {
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Start ws服务启动并注册用户相关内容
func (ws *WsServerController) Start(ctx *gin.Context) {

	wsCoon, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		//socket 服务异常
		Util.Error(ctx, Ecode.ServiceErrorCode.Code, Ecode.ServiceErrorCode.Message)
		return
	}
	defer wsCoon.Close()
	//默认开启加密为了数据传输安全
	wsServer := ws2.NewWs(wsCoon, true)
	//todo::启动socket服务链接之后goroutine启动broadcast广播
	//启动之前注入初始化ws route 路由
	engine := interfaceDepend.NewWsRouterDepend()
	wsEngine := interfaceDepend.InitWsRouter(engine)
	wsServer.Router(wsEngine)
	wsServer.Start()
	wsServer.HandShake()

	select {}
}
