package Ws

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Model/Im"
	"api-skeleton/grateway"
	"fmt"
	"strconv"
	"time"
)

type ChatLogic struct {
}

type ChatResData struct {
	Msg string
}

func (cl *ChatLogic) ToChat(req *grateway.WsMsgReq) (error, Im.ImMsg) {
	requestData := req.Body.Data.(map[string]interface{})
	fmt.Println("请求的数据", requestData)
	toUserId, _ := strconv.Atoi(requestData["toUserId"].(string))

	nowTime := time.Now().Unix()
	ImMsgData := Im.ImMsg{
		ToUserId:   int64(toUserId),
		FromUserId: 1,
		Content:    requestData["msg"].(string),
		SendTime:   nowTime,
	}

	err := Global.ImDB.Create(&ImMsgData).Error
	if err != nil {
		return err, ImMsgData
	}

	return nil, ImMsgData
}
