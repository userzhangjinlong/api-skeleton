package Middleware

import (
	"api-skeleton/app/Global"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/uber/jaeger-client-go"
)

func Tracing() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var newCtx context.Context
		var span opentracing.Span

		spanCtx, err := opentracing.GlobalTracer().Extract(
			opentracing.HTTPHeaders,
			opentracing.HTTPHeadersCarrier(ctx.Request.Header),
		)

		if err != nil {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				Global.Tracer,
				ctx.Request.URL.Path,
			)
		} else {
			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				ctx.Request.Context(),
				Global.Tracer,
				ctx.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
			)
		}

		defer span.Finish()
		ctx.Request = ctx.Request.WithContext(newCtx)

		var traceID string
		var spanID string
		var spanContext = span.Context()
		switch spanContext.(type) {
		case jaeger.SpanContext:
			jaegerContext := spanContext.(jaeger.SpanContext)
			traceID = jaegerContext.TraceID().String()
			spanID = jaegerContext.SpanID().String()
		}

		ctx.Set("X-Trace-ID", traceID)
		ctx.Set("X-Span-ID", spanID)
		ctx.Request = ctx.Request.WithContext(ctx)

		ctx.Next()
	}
}
