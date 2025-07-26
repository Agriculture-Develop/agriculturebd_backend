package logs

import (
	"github.com/Agriculture-Develop/agriculturebd/api/config"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"os"
)

var Logger *zap.SugaredLogger

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
	zLogger := zap.New(core, zap.AddCaller())

	Logger = zLogger.Sugar()
	zap.ReplaceGlobals(zLogger)
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
