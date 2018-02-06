package main

import (
	"runtime"
	"time"

	"ssf4g/common/timer"
	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/server/resource-srv/common/srv-const"
)

// 系统Routine, 定时GC
func SysRoutine() {
	gc_timer := make(chan int32, 10)
	gc_timer <- 1

	for {
		select {
		case <-gc_timer:
			runtime.GC()

			tlog.Info("gc executed.")
			tlog.Info("goroutine cnt info (%d).", runtime.NumGoroutine())
			tlog.Info("gc summary info (%s).", utility.GCSummary())

			timer.Add(0, time.Now().Unix()+srvconst.GC_INTERVAL, gc_timer)
		}
	}
}
