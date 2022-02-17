package Api

import (
	"api-skeleton/app/Cache"
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
	logrus.WithFields(logrus.Fields{
		"code": 200,
		"data": "success",
	}).Info("测试日志写入12")
	Util.Success(ctx, map[string]string{
		"data": val,
	})
}
