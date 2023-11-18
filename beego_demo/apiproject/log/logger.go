package log

import (
	"apiproject/conf"
	"apiproject/constant"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/logs"
	"strings"
)

// LoggerConfig 日志模块, 需要初始化才可以使用
type LoggerConfig struct {
	FileName string `json:"filename"`
}

var logCallDepth int = 3
var beeLogger *logs.BeeLogger

func init() {
	beeLogger = logs.NewLogger()
	var logCfg = LoggerConfig{
		FileName: conf.GetLogFilename(),
	}
	logCfgJson, _ := json.Marshal(&logCfg)
	logAdapter := conf.GetLogAdapter()
	switch logAdapter {
	case constant.AdapterFile:
		// 单个日志文件
		beeLogger.SetLogger(constant.AdapterFile, string(logCfgJson))
	case constant.AdapterConsole:
		// 控制台记录
		beeLogger.SetLogger(constant.AdapterConsole, ``)
	case constant.AdapterConsoleFile:
		// 打印到控制台和文件
		beeLogger.SetLogger(constant.AdapterFile, string(logCfgJson))
		beeLogger.SetLogger(constant.AdapterConsole, ``)
	default:
		beeLogger.SetLogger(constant.AdapterFile, string(logCfgJson))
	}
	//设置输出等级
	beeLogger.SetLevel(conf.GetLogLevel())
	// 输出文件名和行号，默认为true
	beeLogger.EnableFuncCallDepth(conf.GetLogFuncCall())
	// 函数行号打印深度，默认为2
	beeLogger.SetLogFuncCallDepth(logCallDepth)
	// 提升性能, 可以设置异步输出
	beeLogger.Async()
	beeLogger.Info("logger init success")
}

func Debug(template string, args ...interface{}) {
	beeLogger.Debug(formatLog(template, args...))
}

func Info(template string, args ...interface{}) {
	beeLogger.Info(formatLog(template, args...))
}

func Warn(template string, args ...interface{}) {
	beeLogger.Warn(formatLog(template, args...))
}

func Error(template string, args ...interface{}) {
	beeLogger.Error(formatLog(template, args...))
}

func formatLog(f interface{}, v ...interface{}) string {
	var msg string
	switch f.(type) {
	case string:
		msg = f.(string)
		if len(v) == 0 {
			return msg
		}
		if strings.Contains(msg, "%") && !strings.Contains(msg, "%%") {
		} else {
			msg += strings.Repeat(" %v", len(v))
		}
	default:
		msg = fmt.Sprint(f)
		if len(v) == 0 {
			return msg
		}
		msg += strings.Repeat(" %v", len(v))
	}
	var data []any
	for i := 0; i < len(v); i++ {
		data = append(data, any(v[i]))
	}
	return fmt.Sprintf(msg, data...)
}
