package grateway

import (
	"github.com/gorilla/websocket"
)

//WsMsgReq ws请求body统一结构封装
type WsMsgReq struct {
	Body *ReqBody
}

//WsMsgRsp ws请求外部统一透出响应体
type WsMsgRsp struct {
	Body *RspBody
}

//RspBody ws请求响应body体
type RspBody struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//ReqBody ws请求body体
type ReqBody struct {
	Path string      `json:"path"` //路由key
	Data interface{} `json:"data"` //请求的消息json数据
}

//SignalChatMsg 单体消息
type SignalChatMsg struct {
	SignalCoon *websocket.Conn
	Msg        *WsMsgRsp
}

var (
	SignalChatMsgChan = make(chan *SignalChatMsg, 1000)
)
