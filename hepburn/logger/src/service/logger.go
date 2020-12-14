package service

import (
	"encoding/json"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

// error logger
var errorLogger *zap.SugaredLogger

var levelMap = map[string]zapcore.Level{
	"debug":  zapcore.DebugLevel,
	"info":   zapcore.InfoLevel,
	"warn":   zapcore.WarnLevel,
	"error":  zapcore.ErrorLevel,
	"dpanic": zapcore.DPanicLevel,
	"panic":  zapcore.PanicLevel,
	"fatal":  zapcore.FatalLevel,
}

func getLoggerLevel(lvl string) zapcore.Level {
	if level, ok := levelMap[lvl]; ok {
		return level
	}
	return zapcore.InfoLevel
}
//后续对日志进行分组
func initLogger(levelParam string) {
	date := time.Now().Format("2006-01-02")
	fileName := Env("logPath") + date + "zap.log"
	level := getLoggerLevel(levelParam)
	hook := lumberjack.Logger{
		Filename:   fileName, // 日志文件路径
		MaxSize:    1,                      // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 30,                       // 日志文件最多保存多少个备份
		MaxAge:     7,                        // 文件最多保存多少天
		Compress:   false,                     // 是否压缩
		LocalTime: true,
	}

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		//EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05"))
		},
		//EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeDuration: func(d time.Duration, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendInt64(int64(d) / 1000000)
		},
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevelAt(level)
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(&hook)), // 打印到文件
		//zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel,                                                                     // 日志级别
	)
	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.AddCaller()
	// 设置初始化字段
	//filed := zap.Fields(zap.String("serviceName", "serviceName"))
	// 构造日志
	logger := zap.New(core, caller, development)

	//fileName := Env("logPath") + "zap.log"
	//level := getLoggerLevel(levelParam)
	//syncWriter := zapcore.AddSync(&lumberjack.Logger{
	//	Filename:  fileName,
	//	MaxSize:   1 << 30, //1G
	//	LocalTime: true,
	//	Compress:  true,
	//})
	//encoder := zap.NewProductionEncoderConfig()
	//encoder.EncodeTime = zapcore.ISO8601TimeEncoder
	//core := zapcore.NewCore(zapcore.NewJSONEncoder(encoder), syncWriter, zap.NewAtomicLevelAt(level))
	//logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	errorLogger = logger.Sugar()
	defer errorLogger.Sync()
}

func Debug(args ...interface{}) {
	initLogger("debug")
	errorLogger.Debug(args...)
}

func Debugf(template string, args ...interface{}) {
	initLogger("debug")
	errorLogger.Debugf(template, args...)
}

func Info(args ...interface{}) {
	initLogger("info")
	errorLogger.Info(args...)
}

func Infof(template string, args ...interface{}) {
	initLogger("info")
	errorLogger.Infof(template, args...)
}

func Warn(args ...interface{}) {
	initLogger("warn")
	errorLogger.Warn(args...)
}

func Warnf(template string, args ...interface{}) {
	initLogger("warn")
	errorLogger.Warnf(template, args...)
}

func Error(args ...interface{}) {
	initLogger("error")
	errorLogger.Error(args...)
}

func Errorf(template string, args ...interface{}) {
	initLogger("error")
	errorLogger.Errorf(template, args...)
}

func DPanic(args ...interface{}) {
	initLogger("dpanic")
	errorLogger.DPanic(args...)
}

func DPanicf(template string, args ...interface{}) {
	initLogger("dpanic")
	errorLogger.DPanicf(template, args...)
}

func Panic(args ...interface{}) {
	initLogger("panic")
	errorLogger.Panic(args...)
}

func Panicf(template string, args ...interface{}) {
	initLogger("panic")
	errorLogger.Panicf(template, args...)
}

func Fatal(args ...interface{}) {
	initLogger("fatal")
	errorLogger.Fatal(args...)
}

func Fatalf(template string, args ...interface{}) {
	initLogger("fatal")
	errorLogger.Fatalf(template, args...)
}

func SendEmail(url string, mailTo map[string]interface{})  {
	data, _:= json.Marshal(mailTo)
	request, _ := http.NewRequest("POST", url, strings.NewReader(string(data)))
	//post数据并接收http响应
	resp,err :=http.DefaultClient.Do(request)
	if err!=nil{
		fmt.Printf("post data error:%v\n",err)
	}else {
		respBody,_ :=ioutil.ReadAll(resp.Body)
		fmt.Printf("response data:%v\n",string(respBody))
	}
}