package middleware

import (
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"regexp"
)

// ZapLogger returns a gin middleware that logs request in a structured format.
func ZapLogger() gin.HandlerFunc {
	return ginzap.GinzapWithConfig(zap.L(), &ginzap.Config{

		// 跳过路径
		SkipPathRegexps: []*regexp.Regexp{
			regexp.MustCompile(`^/health/?$`),
			regexp.MustCompile(`^/metrics/?$`),
			regexp.MustCompile(`^/ping/?$`),
		},

		// 额外封装信息
		Context: func(c *gin.Context) []zap.Field {
			fields := []zap.Field{
				zap.String("service_name", "gateway"),
			}

			// trace 信息（可选）
			if span := trace.SpanFromContext(c.Request.Context()); span.SpanContext().IsValid() {
				fields = append(fields,
					zap.String("trace_id", span.SpanContext().TraceID().String()),
					zap.String("span_id", span.SpanContext().SpanID().String()),
				)
			}

			return fields
		},
	})
}
