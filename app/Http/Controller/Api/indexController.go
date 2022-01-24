package Api

import "github.com/gin-gonic/gin"

type Index struct {
}

func (i *Index) Index(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": "123124",
	})
}
