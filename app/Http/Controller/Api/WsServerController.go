package Api

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Util"
	"api-skeleton/app/interfaceDepend"
	"api-skeleton/grateway"
	"api-skeleton/ws/Client"
	ws2 "api-skeleton/ws/Server"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
	"strconv"
	"sync"
)

type WsServerController struct {
	Base *BaseController
	Lock sync.RWMutex
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//Start 启动当前客户端socket链接服务并注册对应路由回掉
func (ws *WsServerController) Start(ctx *gin.Context) {
	token, _ := ctx.GetQuery("token")
	userInfo, _ := Util.ParseToken(token)
	userinfoId, err := strconv.Atoi(userInfo.ID)
	//这里的socket coon 只是当前客户端请求链接的一个 不是所有的客户端 应该将所有的客户端保存起来
	wsCoon, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		//socket 服务异常
		Util.Error(ctx, Ecode.ServiceErrorCode.Code, Ecode.ServiceErrorCode.Message)
		return
	}
	defer wsCoon.Close()

	//ws服务请求 秘钥下发
	secretKey := Util.RandSeeks(16)

	//管理当前客户端已经下发的ws服务链接
	socketCoon := &Client.SocketCoon{
		Coon:      wsCoon,
		SecretKey: secretKey,
		KeepAlive: true,
		RspMsg:    make(chan *grateway.WsMsgRsp, 1000),
	}
	ws.Lock.Lock()
	Client.UserSocketConn[userinfoId] = socketCoon
	ws.Lock.Unlock()

	//注册当前socket链接服务并初始化路由 中间件 握手等
	//todo::通过当前用户的链接注册一些服务 路由 握手 中间件等
	wsServer := ws2.NewWs(wsCoon)
	//todo::启动socket服务链接之后goroutine启动broadcast广播
	//启动之前注入初始化ws route 路由
	engine := interfaceDepend.NewWsRouterDepend()
	wsEngine := interfaceDepend.InitWsRouter(engine)
	wsServer.Router(wsEngine)
	//todo::携程消息服务启动
	wsServer.Start()
	//握手成功上报秘钥
	wsServer.HandShake(secretKey)

	select {}
}
