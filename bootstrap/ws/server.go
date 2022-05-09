package ws

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Util"
	"api-skeleton/grateway"
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
type wsServer struct {
	coon         *websocket.Conn         //ws 服务链接
	rspMsg       chan *grateway.WsMsgRsp //响应消息体
	router       *wsRoutes.WsEngine      //ws路由处理
	property     map[string]interface{}  //ws服务配置的一些属性
	propertyLock sync.RWMutex            //属性读写锁
	seq          int64                   //请求数据序号
	isEncrypt    bool                    //是否需要加密默认需要加密
}

//NewWs 构造函数实例化返回ws服务
func NewWs(coon *websocket.Conn, isEncrypt bool) (ws *wsServer) {
	ws = &wsServer{
		coon:      coon,
		rspMsg:    make(chan *grateway.WsMsgRsp, 1000),
		property:  make(map[string]interface{}),
		seq:       0,
		isEncrypt: isEncrypt,
		router:    wsRoutes.NewWsEngine(),
	}

	return
}

//SetProperty 设置服务端属性
func (ws *wsServer) SetProperty(key string, value interface{}) {
	ws.propertyLock.Lock()
	defer ws.propertyLock.Unlock()
	ws.property[key] = value
}

//GetProperty 获取服务属性
func (ws *wsServer) GetProperty(key string) (interface{}, error) {
	ws.propertyLock.RLock()
	defer ws.propertyLock.RUnlock()
	if value, ok := ws.property[key]; ok {
		return value, nil
	} else {
		return nil, errors.New("get property error")
	}
}

//RemoveProperty 移除掉服务属性
func (ws *wsServer) RemoveProperty(key string) {
	ws.propertyLock.Lock()
	defer ws.propertyLock.Unlock()
	delete(ws.property, key)
}

//Push 压进消息给服务端接收处理
func (ws *wsServer) Push(data interface{}) {

	response := &grateway.WsMsgRsp{
		Body: &grateway.RspBody{
			Code: Ecode.ResponseOkCode.Code,
			Msg:  Ecode.ResponseOkCode.Message,
			Data: data,
		},
	}
	ws.rspMsg <- response
}

//Start ws服务启动
func (ws *wsServer) Start() {
	go ws.Read()
	go ws.Write()
}

//Read 读取客户端发送过来的ws数据
func (ws *wsServer) Read() {
	for {
		//读取ws中的数据
		_, message, err := ws.coon.ReadMessage()
		fmt.Printf("读取到的数据message:%v", string(message))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": message,
				"err": err,
			}).Error("读取ws消息异常")
		}
		reqData := new(grateway.WsMsgReq)
		//todo::解析读取到的json消息并且执行对应路由返回响应
		err = json.Unmarshal(message, &reqData.Body)
		if err != nil {
			//返回响应请求的数据格式错误
			panic(err)
		}

		ws.router.Run(reqData.Body.Path, reqData)
		ws.Push(reqData.Body.Path)
	}
}

//Write 获取到客户端ws消息响应服务端ws消息
func (ws *wsServer) Write() {
	for {
		message := <-ws.rspMsg
		fmt.Printf("写入的ws数据:%v", message.Body.Data)
		//响应给客户端ws数据
		data, err := json.Marshal(message.Body)
		if err != nil {
			log.Println(err)
		}
		err = ws.coon.WriteMessage(websocket.BinaryMessage, data)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": message,
				"err": err,
			}).Error("写入ws消息异常")
		}
	}
}

//Router ws服务路由注入
func (ws *wsServer) Router(router *wsRoutes.WsEngine) {
	ws.router = router
}

//HandShake ws握手响应秘钥给客户端
func (ws *wsServer) HandShake() {
	secretKey := ""
	val, err := ws.GetProperty(secretKey)
	if err != nil {
		//设置secretKey
		secretKey = Util.RandSeeks(16)
	} else {
		secretKey = val.(string)
	}

	//封装请求 响应握手秘钥
	req := grateway.WsMsgReq{
		Body: &grateway.ReqBody{
			Path: secretKey,
			Data: "ws.handShake",
		},
	}
	ws.Push(req)
}
