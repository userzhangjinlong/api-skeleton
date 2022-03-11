package Api

import (
	"api-skeleton/app/Cache"
	"api-skeleton/app/Global"
	"api-skeleton/app/Model/ApiSkeleton"
	"api-skeleton/app/Util"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Index struct {
}

//Index 首页入口
func (i *Index) Index(ctx *gin.Context) {
	//入参校验
	//param := struct {
	//	Name  string `form:"name" binding:"required,max=2,min=1"`
	//	State int8   `form:"state,default=1" binding:"oneof=0 1"`
	//}{}
	//valid, errs := Request.BindAndValid(ctx, &param)
	//if !valid {
	//	errString := fmt.Sprintf("参数错误：%s", errs)
	//	Util.Error(ctx, 400, errString)
	//	return
	//}
	var cache Cache.BaseRedis
	val, _ := cache.Get("test")
	//集群使用获取方式
	//val := Global.RedisCluster.Get("test").Val()
	//val, err1 := cache.Get("test")
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	traceId, _ := ctx.Get("X-Trace-ID")
	spanId, _ := ctx.Get("X-Span-ID")
	userinfo, _ := ctx.Get("User")
	user := ApiSkeleton.User{}
	Global.DB.Where("tel = ?", "18030769533").Find(&user)
	logrus.WithFields(logrus.Fields{
		"code":     200,
		"data":     "success",
		"trace_id": traceId,
		"span_id":  spanId,
	}).Info("测试日志写入12")
	result := struct {
		Val       string      `json:"vals"`
		LoginInfo interface{} `json:"userinfo"`
		User      interface{} `json:"user"`
		Data      interface{} `json:"data"`
	}{}

	result.Val = val
	result.LoginInfo = userinfo
	result.User = user

	//curl 客户端工具调试
	getAddress := "http://qa.wpt.la/mofei/japi/user/findUser?verifyStatus=individual_verify_pass&size=10&phone=true&scene=weertre"
	header := make(map[string]string, 1)
	header["Cookie"] = "wpt_env_num=test-06"
	resData := Util.CurlRequestGet(getAddress, header, nil)
	res := struct {
		Code int `json:"code"`
		Data struct {
			Code int `json:"code"`
			Data []struct {
				ID    string `json:"id"`
				Phone string `json:"phone"`
				URI   string `json:"uri"`
			} `json:"data"`
			Msg string `json:"msg"`
		} `json:"data"`
		Msg string `json:"msg"`
	}{}
	json.Unmarshal(resData, &res)
	//消息投递 创建数据
	//测试写入nsq消息
	//生产者写入nsq,10条消息，topic = "test"
	topic := "createRankingMessage"
	topic2 := "createRankingMessageNode2"
	topic3 := "createRankingMessageNode3"
	for _, v := range res.Data.Data {
		Util.DeliveryNsq1Message(topic, []byte(v.ID))
		Util.DeliveryNsq2Message(topic2, []byte(v.ID))
		Util.DeliveryNsq3Message(topic3, []byte(v.ID))
	}
	result.Data = res
	Util.Success(ctx, result)
}
