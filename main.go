package main

import (
	. "comment/config"
	"comment/logger"
	"comment/middleware"
	"comment/router"
	"comment/util"
	"github.com/gin-gonic/gin"
	"strings"
)

func main() {
	//设置日志前台可见
	consoleFlag := false
	//set log level and init logger
	if Log.IsConsole == "1" {
		consoleFlag = true
	}
	logger.SetConsole(consoleFlag)

	// 初始化日志等级
	logLevel := logger.ERROR
	if strings.EqualFold(Log.Level, "debug") {
		logLevel = logger.DEBUG
	} else if strings.EqualFold(Log.Level, "info") {
		logLevel = logger.INFO
	} else if strings.EqualFold(Log.Level, "warn") {
		logLevel = logger.WARN
	} else if strings.EqualFold(Log.Level, "error") {
		logLevel = logger.ERROR
	} else if strings.EqualFold(Log.Level, "fatal") {
		logLevel = logger.FATAL
	}
	logger.SetLevel(logLevel)
	//根据配置文件，设置日志路径，日志名，日志切割大小限制
	logger.SetRollingFile(Log.Path, "server.log", 10, 10, logger.MB)
	// 初始化 gin 服务
	r := gin.Default()
	// 初始化中间件
	middleware.Init(r)
	// 初始化路由
	router.Init(r)
	r.Run(util.Join(util.S{Global.Host, Global.Port}, ":"))
}
