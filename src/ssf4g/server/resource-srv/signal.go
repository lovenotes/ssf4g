package main

import (
	"os"
	"os/signal"
	"syscall"

	"ssf4g/common/config"
	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/server/resource-srv/common/resource-data"
	"ssf4g/server/resource-srv/common/srv-config"
)

// Func - 信号量处理
func SignalProc() {
	ch := make(chan os.Signal, 1)

	signal.Notify(ch, syscall.SIGHUP, syscall.SIGABRT, syscall.SIGTERM, syscall.SIGKILL, syscall.SIGINT)

	defer func() {
		if x := recover(); x != nil {
			tlog.Error("signal proc err (panic %v).", x)
		}

		utility.RemovePIDFile()

		os.Exit(0)
	}()

	for {
		msg := <-ch

		switch msg {
		case syscall.SIGABRT:
			utility.RemovePIDFile()

			panic("syscall.SIGABRT")

			break
		case syscall.SIGHUP:
			tlog.Info("[SIGHUP]")

			// 重新加载Config文件
			config.ReloadConfig()
			srvconfig.ReloadSrvConfig()

			// 重新加载ResourceData
			resourcedata.ReloadResourceData()

			break
		case syscall.SIGINT:
			tlog.Info("[SIGINT]")

			break
		case syscall.SIGTERM:
			tlog.Info("[SIGTERM]")

			utility.RemovePIDFile()

			os.Exit(0)

			break
		case syscall.SIGKILL:
			tlog.Info("[SIGKILL]")

			break
		}
	}
}
