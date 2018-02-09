package redismgr

import (
	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/common/srv-config"
	"ssf4g/server/login-srv/redis/account-redis"
)

var (
	_account_redis *accountredis.AccountRedis
)

func init() {
	maxIdleConn := srvconfig.GetConfig().RedisMaxIdleConn
	timeout := srvconfig.GetConfig().RedisTimeout

	accountUrl := srvconfig.GetConfig().AccountRedisUrl
	accountAuth := srvconfig.GetConfig().AccountRedisAuth

	_account_redis = &accountredis.AccountRedis{}

	errData := _account_redis.InitAccountRedis(accountUrl, maxIdleConn, timeout, accountAuth)

	if errData != nil {
		errMsg := tlog.Error("init account redis (%s, %d, %d, %s) err (%v).", accountUrl, maxIdleConn, timeout, accountAuth, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	return
}

func GetAccountRedis() *accountredis.AccountRedis {
	return _account_redis
}
