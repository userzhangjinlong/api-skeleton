package Middleware

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"strings"
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
		param, _ := GetFormParam(ctx)
		fields := logrus.Fields{
			"request":  param,
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

// GetFormParam 获取form表单请求参数和url参数
//raw仅支持格式foo=bar\n baz=The first line类数据
func GetFormParam(ctx *gin.Context) (data map[string]string, err error) {
	contextType := ctx.Request.Header.Get("Content-Type")
	data = make(map[string]string)

	if strings.Contains(contextType, "multipart/form-data") ||
		strings.Contains(contextType, "application/x-www-form-urlencoded") {
		//解析Form
		ctx.Request.ParseMultipartForm(128)
		//说明:post方法,'Content-Type': 'application/x-www-form-urlencoded'和multipart/form-data
		for key, valueArray := range ctx.Request.Form {
			if len(valueArray) > 1 {
				errMsg := fmt.Sprintf("#ERROR#[%s]参数设置了[%d]次,只能设置一次.", key, len(valueArray))
				return nil, errors.New(errMsg)
			}
			data[key] = ctx.PostForm(key)
		}
	} else if strings.Contains(contextType, "text/plain") {
		bodyBytes, err := ctx.GetRawData()
		if err == nil {
			param := string(bodyBytes)
			for _, v := range strings.Split(param, "\n") {
				paramData := strings.Split(v, "=")
				if paramData != nil &&
					paramData[0] != "" &&
					paramData[1] != "" {
					data[paramData[0]] = paramData[1]
				}
			}

		}
	} else {
		//todo:binary 文件流请求 文件流请求建议走七牛 阿里等三方传地址过来暂不兼容文件流日志记录

	}

	//获取url参数
	for key, _ := range ctx.Request.URL.Query() {
		data[key] = ctx.Query(key)
	}

	return
}
