package rpcservice

import (
	"net"

	"ssf4g/common/tlog"
	"ssf4g/protobuf/game-proto"
	"ssf4g/server/game-srv/common/srv-config"
	"ssf4g/server/game-srv/rpc-service/portal-service/handler"

	"google.golang.org/grpc"
)

func StartPortalRpcService() {
	serviceRPC := srvconfig.GetConfig().ServiceRPC

	lis, err := net.Listen("tcp", serviceRPC)

	if err != nil {
		errMsg := tlog.Error("start rpc service (%s, %s) err (listen %v).", srvconfig.GetConfig().SrvName, serviceRPC, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}

	server := grpc.NewServer()

	// 注册RPC-Service
	gameproto.RegisterPortalIntrServiceServer(server, &portalhandler.PortalIntrService{})

	err = server.Serve(lis)

	if err != nil {
		errMsg := tlog.Error("start rpc service (%s, %s) err (%v).", srvconfig.GetConfig().SrvName, serviceRPC, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}
}
