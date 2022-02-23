package Api

import (
	"api-skeleton/app/Cache"
	"api-skeleton/app/Global"
	"api-skeleton/app/Model"
	"api-skeleton/app/Util"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Index struct {
}

//Index 首页入口
func (i *Index) Index(ctx *gin.Context) {
	//入参校验
	//param := struct {
	//	Name  string `form:"name" binding:"required,max=1"`
	//	State int8   `form:"state,default=1" binding:"oneof=0 1"`
	//}{}
	//valid, errs := Request.BindAndValid(ctx, &param)
	//if !valid {
	//	fmt.Printf("参数错误：%s", errs)
	//	return
	//}
	var cache Cache.BaseRedis
	val, _ := cache.HGet("user", "1")
	//val, err1 := cache.Get("test")
	//if err1 != nil {
	//	fmt.Println(err1)
	//}
	traceId, _ := ctx.Get("X-Trace-ID")
	spanId, _ := ctx.Get("X-Span-ID")
	userinfo, _ := ctx.Get("User")
	user := Model.User{}
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
	}{}

	result.Val = val
	result.LoginInfo = userinfo
	result.User = user
	Util.Success(ctx, result)
}
