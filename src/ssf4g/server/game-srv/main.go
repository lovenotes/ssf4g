package main

import (
	"runtime"

	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/server/game-srv/common/srv-config"
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

	// 初始化Config信息
	srvconfig.InitSrvConfig()
}
