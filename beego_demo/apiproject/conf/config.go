package conf

import (
	"apiproject/constant"
	"github.com/astaxie/beego"
)

func GetDbName() string {
	return beego.AppConfig.String("mysqlName")
}

func GetDbPort() int {
	port, _ := beego.AppConfig.Int("mysqlPort")
	return port
}

func GetDbHost() string {
	return beego.AppConfig.String("mysqlHost")
}

func GetDbUser() string {
	return beego.AppConfig.String("mysqlUser")
}

func GetDbPassword() string {
	return beego.AppConfig.String("mysqlPassword")
}

func GetDbMaxConnect() int {
	connect, _ := beego.AppConfig.Int("maxConnect")
	return connect
}

func GetDbMaxOpenConnect() int {
	connect, _ := beego.AppConfig.Int("maxOpenConnect")
	return connect
}

func GetLogAdapter() string {
	return beego.AppConfig.String("logAdapter")
}

func GetLogLevel() int {
	level, _ := beego.AppConfig.Int("logLevel")
	return level
}

func GetLogFuncCall() bool {
	result, _ := beego.AppConfig.Bool("logFuncCall")
	return result
}

func GetLogFilename() string {
	return beego.AppConfig.String("log_filename")
}

func GetRpcIp() string {
	return beego.AppConfig.String("rpcIp")
}

func GetRpcServicePort() int {
	port, _ := beego.AppConfig.Int("rpcServicePort")
	return port
}

func GetRedisHost() string {
	return beego.AppConfig.String("redisHost")
}

func GetRedisPort() int {
	port, _ := beego.AppConfig.Int("redisPort")
	return port
}

func GetRedisPassword() string {
	return beego.AppConfig.String("redisPassword")
}

func GetRedisDb() int {
	db, _ := beego.AppConfig.Int("redisDb")
	return db
}

func GetRedisMaxIdle() int {
	idle, _ := beego.AppConfig.Int("redisMaxIdle")
	return idle
}

func GetRedisMaxActive() int {
	active, _ := beego.AppConfig.Int("redisMaxActive")
	return active
}

func GetRedisIdleTimeout() int {
	time, _ := beego.AppConfig.Int("redisIdleTimeout")
	return time
}

func GetWebSocketPath() string {
	return beego.AppConfig.String("wsPath")
}

func GetWebSocketPort() int {
	port, _ := beego.AppConfig.Int("wsPort")
	return port
}

func GetPprofPort() int {
	port, _ := beego.AppConfig.Int("pprofPort")
	return port
}

func GetPositionOffTime() int64 {
	offTime, err := beego.AppConfig.Int64("positionOfflineTime")
	if err != nil || offTime <= constant.ParamsDefault {
		offTime = constant.TagOfflineTime
	}
	return offTime
}
