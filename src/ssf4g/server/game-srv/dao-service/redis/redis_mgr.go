package redismgr

import (
	"ssf4g/common/tlog"
	"ssf4g/server/game-srv/common/srv-config"
	"ssf4g/server/game-srv/dao-service/redis/server-redis"
)

var (
	_server_redis *serverredis.ServerRedis
)

func init() {
	maxIdleConn := srvconfig.GetConfig().RedisMaxIdleConn
	timeout := srvconfig.GetConfig().RedisTimeout

	serverUrl := srvconfig.GetConfig().ServerRedisUrl
	serverAuth := srvconfig.GetConfig().ServerRedisAuth

	_server_redis = &serverredis.ServerRedis{}

	errData := _server_redis.InitServerRedis(serverUrl, maxIdleConn, timeout, serverAuth)

	if errData != nil {
		errMsg := tlog.Error("init account redis (%s, %d, %d, %s) err (%v).", serverUrl, maxIdleConn, timeout, serverAuth, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	return
}

func GetServerRedis() *serverredis.ServerRedis {
	return _server_redis
}
