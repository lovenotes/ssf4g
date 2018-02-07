package main

import (
	"runtime"

	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/server/resource-srv/common/resource-data"
	"ssf4g/server/resource-srv/common/srv-config"
	"ssf4g/server/resource-srv/service/http-service"
)

func main() {
	defer func() {
		if x := recover(); x != nil {
			tlog.Error("caught panic in main(%v)", x)
		}
	}()

	// 生成PID文件
	utility.GenPIDFile()

	// 设置CPU最大数量
	runtime.GOMAXPROCS(runtime.NumCPU())

	// 初始化Srv Config信息
	srvconfig.InitSrvConfig()

	// 启动Signal监听及性能监听Goroutine
	startup()

	// 初始化游戏逻辑相关的各个模块
	initModel()

	// 启动GM Service
	go httpservice.StartGmService()

	// 启动Http Service
	httpservice.StartHttpService()
}

// 启动其他相关Routine
func startup() {
	// 捕获并处理UNIX信号
	go SignalProc()

	// 性能监测GoRoutine
	go SysRoutine()
}

// 初始化各个模块
func initModel() {
	// 初始化ResourceData
	resourcedata.InitResourceData()
}
