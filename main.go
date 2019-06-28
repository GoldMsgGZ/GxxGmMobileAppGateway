package main

import (
	"fmt"
	"github.com/gogf/gf/g"
	"github.com/gogf/gf/g/os/glog"
)

func main() {
	fmt.Println("GxxGmMobileAppGatreway is running ...")

	// 配置对象
	config := g.Config()

	// 日志系统
	logpath := config.GetString("logger.path")
	glog.SetPath(logpath)
	glog.SetStdoutPrint(true)

	// http服务
	server := g.Server(config.GetString("setting.http_name"))

	server.SetLogPath(logpath)
	server.SetErrorLogEnabled(true)
	server.SetAccessLogEnabled(true)
	server.SetPort(config.GetInt("setting.http_port"))

	server.Start()

	g.Wait()
}
