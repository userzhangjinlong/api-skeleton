package Util

import (
	"bytes"
	"fmt"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go/config"
	"io"
	"runtime"
	"time"
)

//NewJaegerTracer 初始化jaeger追中配置
func NewJaegerTracer(serviceName, agentHostPort string) (opentracing.Tracer, io.Closer, error) {
	cfg := &config.Configuration{
		ServiceName: serviceName,
		Sampler: &config.SamplerConfig{
			Type:  "const",
			Param: 1,
		},
		Reporter: &config.ReporterConfig{
			LogSpans:            true,
			BufferFlushInterval: 1 * time.Second,
			LocalAgentHostPort:  agentHostPort,
		},
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		return nil, nil, err
	}
	opentracing.SetGlobalTracer(tracer)
	return tracer, closer, err
}

//PanicTrace 系统级异常调用栈追踪
func PanicTrace(err interface{}) string {
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "%v\r\n", err)
	for i := 1; ; i++ {
		pc, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		fmt.Fprintf(buf, "%s,%d(0x%x)\r\n", file, line, pc)
	}
	return buf.String()
}
