package memcachedmgr

import (
	"ssf4g/common/memcached"
	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/common/srv-config"
	"ssf4g/server/login-srv/memcached/account-memecached"
)

var (
	_account_memcached *accountmemcached.AccountMemcached
)

func init() {
	maxOpenConn := srvconfig.GetConfig().MemcachedMaxOpenConn
	memcachedUrl := srvconfig.GetConfig().MemcachedUrl

	memcached, err := memcached.NewMemcached(memcachedUrl, maxOpenConn)

	if err != nil || memcached == nil {
		if err != nil {
			tlog.Error("init memcached (%d, %d) err (%v).", memcachedUrl, maxOpenConn, err)
		} else {
			tlog.Error("init memcached (%d, %d) err (memcached nil).", memcachedUrl, maxOpenConn)
		}

		return
	}

	_account_memcached = accountmemcached.NewAccountMemcached(memcached)

	return
}

func GetAccountMemcached() *accountmemcached.AccountMemcached {
	return _account_memcached
}
