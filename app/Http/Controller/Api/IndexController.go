package Api

import "github.com/gin-gonic/gin"

type Index struct {
}

//Index 首页入口
func (i *Index) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "123124",
	})
}
