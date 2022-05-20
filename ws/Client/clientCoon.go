package Client

import (
	"api-skeleton/grateway"
	"github.com/gorilla/websocket"
)

type SocketCoon struct {
	Coon      *websocket.Conn         //当前链接
	SecretKey string                  //当前链接握手秘钥
	KeepAlive bool                    //链接状态
	RspMsg    chan *grateway.WsMsgRsp //客户端响应消息结构
}

//UserSocketConn 管理服务端响应用户所有的socketCoon链接
var UserSocketConn = make(map[int]*SocketCoon)

//SendMsg 给指定客户端响应消息
func SendMsg(userId int, responseMsg *grateway.WsMsgRsp) {
	if socketClass, ok := UserSocketConn[userId]; ok {
		signalChatMsg := &grateway.SignalChatMsg{
			SignalCoon: socketClass.Coon,
			Msg:        responseMsg,
		}
		grateway.SignalChatMsgChan <- signalChatMsg
	} else {
		//todo::链接不存在时离线消息发送
	}
}
