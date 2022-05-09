package Im

import (
	"api-skeleton/grateway"
	"fmt"
)

type ChatController struct {
}

func (c *ChatController) HeartCheck(req *grateway.WsMsgReq) {
	fmt.Println("是否执行了心跳这里")
	fmt.Println(req)
}

func (c *ChatController) ToChat(req *grateway.WsMsgReq) {
	fmt.Println("是否执行了聊天这里")
	fmt.Println(req.Body.Data)
}
