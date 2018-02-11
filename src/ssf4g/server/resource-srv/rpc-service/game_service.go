package rpcservice

import (
	"net"

	"ssf4g/common/tlog"
	"ssf4g/protobuf/resource-proto"
	"ssf4g/server/resource-srv/common/srv-config"
	"ssf4g/server/resource-srv/rpc-service/game-service/handler"

	"google.golang.org/grpc"
)

func StartGameRpcService() {
	serviceRPC := srvconfig.GetConfig().ServiceRPC

	lis, err := net.Listen("tcp", serviceRPC)

	if err != nil {
		errMsg := tlog.Error("start rpc service (%s, %s) err (listen %v).", srvconfig.GetConfig().SrvName, serviceRPC, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}

	server := grpc.NewServer()

	// 注册RPC-Service
	resourceproto.RegisterGameIntrServiceServer(server, &gamehandler.GameIntrService{})

	err = server.Serve(lis)

	if err != nil {
		errMsg := tlog.Error("start rpc service (%s, %s) err (%v).", srvconfig.GetConfig().SrvName, serviceRPC, err)

		tlog.AsyncSend(tlog.NewErrData(err, errMsg))

		return
	}
}
