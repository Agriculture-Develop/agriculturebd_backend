package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

// GinRecovery 用于替换gin框架的Recovery中间件，因为传入参数，再包一层
func GinRecovery(stack bool) gin.HandlerFunc {
	logger := zap.L()
	return func(c *gin.Context) {
		defer func() {
			// defer 延迟调用，出了异常，处理并恢复异常，记录日志
			if err := recover(); err != nil {
				// 检查 errs 是否是 errs 类型，如果不是，转换为 errs 类型
				var e error
				if er, ok := err.(error); ok {
					e = er
				} else {
					e = fmt.Errorf("%v", err)
				}

				// 检查是否存在断开的连接(broken pipe或者connection reset by peer)
				var brokenPipe bool
				var ne *net.OpError
				if errors.As(e, &ne) {
					var se *os.SyscallError
					if errors.As(ne.Err, &se) {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				// httputil包预先准备好的DumpRequest方法
				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(c.Request.URL.Path,
						zap.Any("errs", e),
						zap.String("request", string(httpRequest)),
					)
					// 如果连接已断开，我们无法向其写入状态
					c.Error(e)
					c.Abort()
					return
				}

				// 是否打印堆栈信息，使用的是debug.Stack()，传入false，在日志中就没有堆栈信息
				if stack {
					logger.Error("[Recovery from panic]",
						zap.Any("errs", e),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())),
					)
				} else {
					logger.Error("[Recovery from panic]",
						zap.Any("errs", e),
						zap.String("request", string(httpRequest)),
					)
				}

				// 有错误，直接返回给前端错误，前端直接报错
				// c.AbortWithStatus(http.StatusInternalServerError)
				// 该方式前端不报错
				c.JSON(200, gin.H{"apiCode": 500, "msg": "服务器异常"})
			}
		}()
		c.Next()
	}
}
