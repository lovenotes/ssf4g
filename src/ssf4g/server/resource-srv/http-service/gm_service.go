package httpservice

import (
	"net/http"
	"time"

	"ssf4g/common/tlog"
	"ssf4g/server/resource-srv/common/srv-config"
	"ssf4g/server/resource-srv/router/gm-router"

	"github.com/gorilla/mux"
)

func StartGmHttpService() {
	// 初始化Router
	muxRouter := mux.NewRouter()

	gmrouter.InitGMRouter(muxRouter)

	// 监听Client的连接
	serviceGM := srvconfig.GetConfig().ServiceGM

	server := &http.Server{
		Handler:      muxRouter,
		Addr:         serviceGM,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		errMsg := tlog.Error("start gm service (%s) err (%v).", srvconfig.GetConfig().SrvName, serviceGM, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}
}
