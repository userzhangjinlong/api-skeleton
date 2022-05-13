package Im

import (
	"api-skeleton/grateway"
	"fmt"
)

type ChatController struct {
}

func (c *ChatController) HeartCheck(req *grateway.WsMsgReq, rsp *grateway.WsMsgRsp) {
	//通过req解析响应rsp
	rsp.Body.Code = 200
	rsp.Body.Msg = "success"
	rsp.Body.Data = "这是我回复的消息"
	fmt.Println(req)
	return
}

func (c *ChatController) ToChat(req *grateway.WsMsgReq, rsp *grateway.WsMsgRsp) {
	fmt.Println("是否执行了聊天这里")
	fmt.Println(req.Body.Data)
}
