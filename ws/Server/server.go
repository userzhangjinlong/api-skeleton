package ws

import (
	"api-skeleton/grateway"
	"api-skeleton/ws/Client"
	"api-skeleton/wsRoutes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"log"
	"sync"
)

//WsServer ws服务端
type WsServer struct {
	coon         *websocket.Conn         //ws服务链接
	router       *wsRoutes.WsEngine      //ws路由处理
	property     map[string]interface{}  //ws服务配置的一些属性
	propertyLock sync.RWMutex            //属性读写锁
	seq          int64                   //请求数据序号
	RspMsg       chan *grateway.WsMsgRsp //服务端响应消息结构
}

var ClientSocketCoon Client.SocketCoon

//NewWs 构造函数实例化返回ws服务
func NewWs(coon *websocket.Conn) (ws *WsServer) {
	ws = &WsServer{
		coon:     coon,
		property: make(map[string]interface{}),
		seq:      0,
		router:   wsRoutes.NewWsEngine(),
		RspMsg:   make(chan *grateway.WsMsgRsp, 1000),
	}

	return
}

//SetProperty 设置服务端属性
func (ws *WsServer) SetProperty(key string, value interface{}) {
	ws.propertyLock.Lock()
	defer ws.propertyLock.Unlock()
	ws.property[key] = value
}

//GetProperty 获取服务属性
func (ws *WsServer) GetProperty(key string) (interface{}, error) {
	ws.propertyLock.RLock()
	defer ws.propertyLock.RUnlock()
	if value, ok := ws.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("get property error")
	}
}

//RemoveProperty 移除掉服务属性
func (ws *WsServer) RemoveProperty(key string) {
	ws.propertyLock.Lock()
	defer ws.propertyLock.Unlock()
	delete(ws.property, key)
}

//Push 服务端主动推送消息
func (ws *WsServer) Push(responseMsg *grateway.WsMsgRsp) {
	ws.RspMsg <- responseMsg
}

//Start ws服务启动
func (ws *WsServer) Start() {
	//从当前socket链接读取消息
	go ws.Read()
	//广播读取到的消息的响应
	go ws.Write()
}

//Read 读取客户端发送过来的ws数据
func (ws *WsServer) Read() {
	for {
		reqData := &grateway.WsMsgReq{
			Body: &grateway.ReqBody{},
		}
		rspData := &grateway.WsMsgRsp{
			Body: &grateway.RspBody{},
		}
		//读取read消息异常捕获抛出
		//todo:: 捕获异常正常抛出之后 之后的消息不会再有响应待处理
		//defer func() {
		//	if err := recover(); err != nil {
		//		rspData.Body.Code = 500
		//		rspData.Body.Msg = "socket服务异常"
		//		rspData.Body.Data = nil
		//		go ws.Push(rspData)
		//	}
		//}()
		//读取ws中的数据
		_, message, err := ws.coon.ReadMessage()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": message,
				"err": err,
			}).Error("读取ws消息异常")
			//服务端读取不到ws链接时关闭
			ws.coon.Close()
			return
		}

		//todo::解析读取到的json消息并且执行对应路由返回响应
		err = json.Unmarshal(message, &reqData.Body)
		if err != nil {
			//返回响应请求的数据格式错误
			//panic(err)
			rspData.Body.Code = 500
			rspData.Body.Msg = "socket服务异常"
			rspData.Body.Data = nil
		}

		if rspData != nil && rspData.Body.Code != 500 {
			ws.router.Run(reqData.Body.Path, reqData, rspData)
		}
		//当前请求路由请求正常响应
		fmt.Println(rspData, "这里收到的消息")
		ws.Push(rspData)

	}
	ws.coon.Close()
}

//Write 获取到客户端ws消息响应服务端ws消息
func (ws *WsServer) Write() {
	for {
		select {
		//服务端响应当前客户端链接请求
		case message := <-ws.RspMsg:
			//响应给当前链接客户端消息
			data, err := json.Marshal(message.Body)
			if err != nil {
				fmt.Println("data json解析异常")
				log.Println(err)
			}
			err = ws.coon.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"msg": message,
					"err": err,
				}).Error("写入ws消息异常")
				//写入消息异常关闭
				ws.coon.Close()
			}
		case signalChatMsg := <-grateway.SignalChatMsgChan:
			//响应给其他客户端链接消息
			data, err := json.Marshal(signalChatMsg.Msg.Body)
			if err != nil {
				fmt.Println("data json解析异常")
				log.Println(err)
			}
			err = signalChatMsg.SignalCoon.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				logrus.WithFields(logrus.Fields{
					"msg": signalChatMsg.Msg,
					"err": err,
				}).Error("写入ws消息异常")
				//写入消息异常关闭
				signalChatMsg.SignalCoon.Close()
			}
		}
		//err = ws.coon.WriteMessage(websocket.TextMessage, data)
		//if err != nil {
		//	fmt.Println("写入ws消息异常")
		//	logrus.WithFields(logrus.Fields{
		//		"msg": message,
		//		"err": err,
		//	}).Error("写入ws消息异常")
		//	//写入消息异常关闭
		//	ws.coon.Close()
		//}

	}
}

//Router ws服务路由注入
func (ws *WsServer) Router(router *wsRoutes.WsEngine) {
	ws.router = router
}

//HandShake ws握手响应秘钥给客户端
func (ws *WsServer) HandShake(secretKey string) {
	//secretKey := ""
	//val, err := ws.GetProperty("secret")
	//if err != nil {
	//	//设置secretKey
	//	secretKey = Util.RandSeeks(16)
	//} else {
	//	secretKey = val.(string)
	//}

	//封装请求 响应握手秘钥 ws握手成功 链接通道建立
	rspData := &grateway.WsMsgRsp{
		Body: &grateway.RspBody{
			200, "success", secretKey,
		},
	}
	ws.SetProperty("secret", secretKey)
	ws.Push(rspData)
}
