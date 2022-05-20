package Ws

import (
	"api-skeleton/app/Global"
	"api-skeleton/app/Model/Im"
	"api-skeleton/app/Util"
	"api-skeleton/grateway"
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
	toUserId, _ := strconv.Atoi(requestData["toUserId"].(string))
	userInfo, err := Util.ParseToken(requestData["token"].(string))
	chatId, _ := strconv.Atoi(userInfo.ID)

	nowTime := time.Now().Unix()
	ImMsgData := Im.ImMsg{
		ToUserId:   int64(toUserId),
		FromUserId: int64(chatId),
		Content:    requestData["msg"].(string),
		SendTime:   nowTime,
	}

	err = Global.ImDB.Create(&ImMsgData).Error
	if err != nil {
		return err, ImMsgData
	}

	return nil, ImMsgData
}
