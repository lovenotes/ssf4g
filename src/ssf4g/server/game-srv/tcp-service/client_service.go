package tcpservice

import (
	"net/http"
	"time"

	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/common/srv-config"
	"ssf4g/server/login-srv/http-service/client-service/router"

	"github.com/gorilla/mux"
)

func StartClientTcpService() {
	// 监听Client的连接
	service := srvconfig.GetConfig().Service

	server := &http.Server{
		Handler:      muxRouter,
		Addr:         service,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := server.ListenAndServe()

	if err != nil {
		errMsg := tlog.Error("start http service (%s, %s) err (%v).", srvconfig.GetConfig().SrvName, service, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}
}
