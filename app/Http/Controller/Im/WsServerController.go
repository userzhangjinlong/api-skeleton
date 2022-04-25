package Im

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Util"
	ws2 "api-skeleton/bootstrap/ws"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

type WsServer struct {
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Start ws服务启动
func (ws *WsServer) Start(ctx *gin.Context) {

	wsCoon, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		//socket 服务异常
		Util.Error(ctx, Ecode.ServiceErrorCode.Code, Ecode.ServiceErrorCode.Message)
		return
	}
	defer wsCoon.Close()
	wsClient := ws2.NewWs(wsCoon)
	//todo::启动socket服务链接之后goroutine启动broadcast广播
	go wsClient.Read()
	go wsClient.Write()

	select {}

}
