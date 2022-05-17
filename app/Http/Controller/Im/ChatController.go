package Im

import (
	"api-skeleton/app/Ecode"
	"api-skeleton/app/Logic/Ws"
	"api-skeleton/app/Model/Im"
	"api-skeleton/grateway"
	"fmt"
)

type ChatController struct {
	chatLogic Ws.ChatLogic
}

func (c *ChatController) HeartCheck(req *grateway.WsMsgReq, rsp *grateway.WsMsgRsp) {
	//通过req解析响应rsp
	rsp.Body.Code = 200
	rsp.Body.Msg = "success"
	rsp.Body.Data = "这是我回复的消息"
	fmt.Println(req)
	return
}

//ToChat 一对一聊天
func (c *ChatController) ToChat(req *grateway.WsMsgReq, rsp *grateway.WsMsgRsp) {
	//保存聊天记录
	var ImMsgModel Im.ImMsg
	err, ImMsgModel := c.chatLogic.ToChat(req)
	if err != nil {
		rsp.Body.Code = Ecode.FailedCode.Code
		rsp.Body.Msg = Ecode.FailedCode.Message
		rsp.Body.Data = err
		return
	}

	rsp.Body.Code = Ecode.ResponseOkCode.Code
	rsp.Body.Msg = Ecode.ResponseOkCode.Message
	rsp.Body.Data = ImMsgModel
	return
}
