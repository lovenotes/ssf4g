package rpcservice

import (
	"net"

	"ssf4g/common/tlog"
	"ssf4g/server/resource-srv/common/srv-config"

	"google.golang.org/grpc"
)

func StartRpcService() {
	serviceRPC := srvconfig.GetConfig().ServiceRPC

	lis, err := net.Listen("tcp", serviceRPC)

	if err != nil {
		errMsg := tlog.Error("start rpc service (%s, %s) err (listen %v).", srvconfig.GetConfig().SrvName, serviceRPC, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}

	server := grpc.NewServer()

	// 注册RPC-Service
	//appproto.RegisterAppIntrServiceServer(srv, &rpcservice.AppIntrService{})

	err = server.Serve(lis)

	if err != nil {
		errMsg := tlog.Error("start rpc service (%s, %s) err (%v).", srvconfig.GetConfig().SrvName, serviceRPC, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}
}
