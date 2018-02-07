package main

import (
	"net/http"
	"runtime"
	"time"

	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/server/resource-srv/common/resource-data"
	"ssf4g/server/resource-srv/common/srv-config"
	"ssf4g/server/resource-srv/router/client-router"
	"ssf4g/server/resource-srv/service/http-service"

	"github.com/gorilla/mux"
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

	// 初始化Router
	muxRouter := mux.NewRouter()

	clientrouter.InitClientRouter(muxRouter)

	// 监听Client的连接
	service := srvconfig.GetConfig().Service

	srv := &http.Server{
		Handler:      muxRouter,
		Addr:         service,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	if err != nil {
		errMsg := tlog.Error("start client service (%s, %s) err (%v).", srvconfig.GetConfig().SrvName, service, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))
	}
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
