package logs

import (
	"errors"
	"fmt"
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/gin-gonic/gin"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func Init() {

	LogConf := config.Get().Log

	// 配置等级
	level := new(zapcore.Level)
	err := level.UnmarshalText([]byte(LogConf.Level))
	if err != nil {
		log.Println(err)
		return
	}

	// 创建Core三大件
	writeSyncer := getLogWriter(LogConf.Filename, LogConf.MaxSize, LogConf.MaxAge, LogConf.MaxBackups)

	encoder := getEncoder()

	// 创建核心-->如果是dev模式，就在控制台和文件都打印，否则就只写到文件中
	var core zapcore.Core

	if LogConf.Mode == "develop" {
		// 开发模式，日志输出到终端
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		// NewTee创建一个核心，将日志条目复制到两个或多个底层核心中。
		core = zapcore.NewTee(
			zapcore.NewCore(encoder, writeSyncer, level),
			zapcore.NewCore(consoleEncoder, zapcore.Lock(os.Stdout), level),
		)
	} else {
		core = zapcore.NewCore(encoder, writeSyncer, level)
	}

	// 创建 logger 对象
	log := zap.New(core, zap.AddCaller())

	// 替换全局的 logger, 后续在其他包中只需使用zap.L()调用即可
	zap.ReplaceGlobals(log)

}

// GinLogger 用于替换gin框架的Logger中间件
func GinLogger(c *gin.Context) {
	logger := zap.L()
	start := time.Now()
	path := c.Request.URL.Path
	query := c.Request.URL.RawQuery
	c.Next() // 执行视图函数
	// 视图函数执行完成，统计时间，记录日志
	cost := time.Since(start)
	logger.Info(path,
		zap.Int("status", c.Writer.Status()),
		zap.String("method", c.Request.Method),
		zap.String("path", path),
		zap.String("query", query),
		zap.String("ip", c.ClientIP()),
		zap.String("user-agent", c.Request.UserAgent()),
		zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
		zap.Duration("cost", cost),
	)
}

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

// 获取Encoder，给初始化logger使用的
func getEncoder() zapcore.Encoder {
	// 使用zap提供的 NewProductionEncoderConfig
	encoderConfig := zap.NewProductionEncoderConfig()
	// 设置时间格式
	encoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")
	// 时间的key
	encoderConfig.TimeKey = "time"
	// 级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	// 显示调用者信息
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}

// 获取切割的问题，给初始化logger使用的
func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
	// 使用 lumberjack 归档切片日志
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberJackLogger)
}
