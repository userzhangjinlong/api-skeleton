package ws

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type WsClient struct {
	Coon    *websocket.Conn
	Message chan []byte
}

func NewWs(coon *websocket.Conn) (ws *WsClient) {
	ws = &WsClient{
		Coon:    coon,
		Message: make(chan []byte),
	}

	return
}

//Read 读取ws数据
func (wc *WsClient) Read() {
	for {
		//读取ws中的数据
		_, message, err := wc.Coon.ReadMessage()
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": message,
				"err": err,
			}).Error("读取ws消息异常")
		}
		wc.Message <- message
	}
}

//Write 写入ws数据
func (wc *WsClient) Write() {
	for {
		message := <-wc.Message
		//写入ws数据
		err := wc.Coon.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"msg": message,
				"err": err,
			}).Error("写入ws消息异常")
		}
	}
}
