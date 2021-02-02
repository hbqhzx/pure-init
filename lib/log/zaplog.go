package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

var zapLog *zap.SugaredLogger

// 格式化时间
func TimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

func InitZapLog(watcher ...string) {
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",                        // json时时间键
		LevelKey:       "level",                       // json时日志等级键
		NameKey:        "logger",                      // json时日志记录器名
		CallerKey:      "caller",                      // json时日志文件信息键
		MessageKey:     "msg",                         // json时日志消息键
		StacktraceKey:  "stacktrace",                  // json时堆栈键
		LineEnding:     zapcore.DefaultLineEnding,     // 友好日志换行符
		EncodeLevel:    zapcore.CapitalLevelEncoder,   // 友好日志等级名大小写（info INFO）
		EncodeTime:     TimeEncoder,                   // 友好日志时日期格式化
		EncodeDuration: zapcore.StringDurationEncoder, // 时间序列化
		EncodeCaller:   zapcore.ShortCallerEncoder,    // 日志文件信息（包/文件.go:行号）
	}

	//zap 不支持文件归档，如果要支持文件按大小或者时间归档，需要使用lumberjack
	apiHook := lumberjack.Logger{
		Filename:   "../log/wukong.log", // 输出文件
		MaxSize:    10,                  // 日志文件最大大小（单位：MB）
		MaxBackups: 5,                   // 保留的旧日志文件最大数量
		MaxAge:     7,                   // 文件最多保存多少天
		LocalTime:  true,                // 使用机器时间，默认false采用utc（美国时间 // ）
	}

	warnHook := lumberjack.Logger{
		Filename:   "../log/warn.log", // 输出文件
		MaxSize:    10,                // 日志文件最大大小（单位：MB）
		MaxBackups: 5,                 // 保留的旧日志文件最大数量
		MaxAge:     7,                 // 文件最多保存多少天
		LocalTime:  true,              // 使用机器时间，默认false采用utc（美国时间 // ）
	}

	watcherHook := lumberjack.Logger{
		Filename:   "../log/watcher.log", // 输出文件
		MaxSize:    10,                   // 日志文件最大大小（单位：MB）
		MaxBackups: 5,                    // 保留的旧日志文件最大数量
		MaxAge:     7,                    // 文件最多保存多少天
		LocalTime:  true,                 // 使用机器时间，默认false采用utc（美国时间 // ）
	}

	apiHighPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.DebugLevel
	})

	warnHighPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.WarnLevel
	})

	// 控制台输出--动态日志等级>=debug
	dynamicLevel := zap.NewAtomicLevelAt(zap.DebugLevel)

	cores := zapcore.NewTee()
	if watcher != nil && watcher[0] == "watcher" {
		cores = zapcore.NewTee(
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&watcherHook), apiHighPriority),
		)
	} else {
		cores = zapcore.NewTee(
			// 友好的格式、输出控制台、动态等级
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), os.Stdout, dynamicLevel),
			// json格式、输出文件、处定义等级规则-->=debug以上级别存入api.log
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&apiHook), apiHighPriority),
			//>=warn以上级别存入warn.log
			zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConfig), zapcore.AddSync(&warnHook), warnHighPriority),
		)
	}

	logger := zap.New(cores, zap.AddCaller(), zap.AddCallerSkip(1))
	zapLog = logger.Sugar()

	zapLog.Info("init log success")
}

//对zaplog进行封装
func Debug(args ...interface{}) {
	zapLog.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	zapLog.Debugf(template, args...)
}

func Info(args ...interface{}) {
	zapLog.Info(args...)
}

func Infof(template string, args ...interface{}) {
	zapLog.Infof(template, args...)
}

func Warn(args ...interface{}) {
	zapLog.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	zapLog.Warnf(template, args...)
}

func Error(args ...interface{}) {
	zapLog.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	zapLog.Errorf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	zapLog.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	zapLog.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	zapLog.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	zapLog.Errorw(msg, keysAndValues...)
}

func Sync() {
	zapLog.Sync()
}

