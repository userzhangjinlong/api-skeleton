package Api

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Index struct {
}

//Index 首页入口
func (i *Index) Index(ctx *gin.Context) {
	logrus.WithFields(logrus.Fields{
		"code": 200,
		"data": "success",
	}).Info("测试日志写入")
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "123124",
	})
}
