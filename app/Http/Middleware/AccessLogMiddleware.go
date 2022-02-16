package Middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"time"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

//Write 双写body 方便取值
func (acs AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := acs.body.Write(p); err != nil {
		return n, err
	}

	return acs.ResponseWriter.Write(p)
}

//AccessLog 访问日志中间件
func AccessLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bodyWriter := &AccessLogWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}

		ctx.Writer = bodyWriter

		beginTime := time.Now().Unix()
		ctx.Next()
		endTime := time.Now().Unix()
		fields := logrus.Fields{
			"request":  ctx.Request.PostForm,
			"response": bodyWriter.body.String(),
		}

		s := "access log: Uri: %s, Host: %s, Ip: %s, Header: %s,method: %s, code: %d, begin_time: %d, end_time: %d"
		logrus.WithFields(fields).Infof(s,
			ctx.Request.RequestURI,
			ctx.Request.Host,
			ctx.ClientIP(),
			ctx.Request.Header,
			ctx.Request.Method,
			bodyWriter.Status(),
			beginTime,
			endTime,
		)
	}
}
