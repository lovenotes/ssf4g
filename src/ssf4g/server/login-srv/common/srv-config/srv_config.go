package srvconfig

import (
	"sync"

	"ssf4g/common/config"
	"ssf4g/common/tlog"
)

const (
	SRV_NAME = "login_srv"

	RUN_MODE = "dev"

	SERVICE     = "0.0.0.0:8001"
	SERVICE_RPC = "0.0.0.0:8021"

	LOG_PATH   = "/data/ssf4g/logs/loginsrv.log"
	SENTRY_DSN = ""

	ACCNT_REGISTER_LIMIT = 102290134

	DB_MAX_IDLE_CONN = 10
	DB_MAX_OPEN_CONN = 100

	REDIS_MAX_IDLE_CONN = 10
	REDIS_TIMEOUT       = 5

	MEMCACHED_MAX_OPEN_CONN = 100

	LOGIN_DB = "ssf4g:ssf4g@(127.0.0.1:3306)/login?timeout=30s&parseTime=true&loc=Local&charset=utf8"

	ACCOUNT_REDIS_URL  = "127.0.0.1:6379"
	ACCOUNT_REDIS_AUTH = ""

	MEMCACHED_URL = "127.0.0.1:11211"
)

type SrvConfig struct {
	SrvName string
	RunMode string

	Service    string
	ServiceRPC string

	LogPath   string
	SentryDsn string

	AccntRegisterLimit uint64

	DBMaxIdleConn int
	DBMaxOpenConn int

	RedisMaxIdleConn int
	RedisTimeout     int

	MemcachedMaxOpenConn int

	LoginDB string

	AccountRedisUrl  string
	AccountRedisAuth string

	MemcachedUrl string
}

var (
	_conf_info *SrvConfig
	_lock      sync.Mutex
)

func InitSrvConfig() {
	_conf_info = &SrvConfig{}

	ReloadSrvConfig()
}

func GetConfig() *SrvConfig {
	return _conf_info
}

func ReloadSrvConfig() {
	_lock.Lock()
	defer _lock.Unlock()

	iniData := config.GetIniData()

	if _conf_info.SrvName = iniData.String("srv_name"); _conf_info.SrvName == "" {
		_conf_info.SrvName = SRV_NAME

		tlog.Warn("reload srv config (%s) warn (default %s).", "srv_name", _conf_info.SrvName)
	}

	if _conf_info.RunMode = iniData.String("run_mode"); _conf_info.RunMode == "" {
		_conf_info.RunMode = RUN_MODE

		tlog.Warn("reload srv config (%s) warn (default %s).", "run_mode", _conf_info.RunMode)
	}

	if _conf_info.Service = iniData.String("service"); _conf_info.Service == "" {
		_conf_info.Service = SERVICE

		tlog.Warn("reload srv config (%s) warn (default %s).", "service", _conf_info.Service)
	}

	if _conf_info.ServiceRPC = iniData.String("service_rpc"); _conf_info.ServiceRPC == "" {
		_conf_info.ServiceRPC = SERVICE_RPC

		tlog.Warn("reload srv config (%s) warn (default %s).", "service_rpc", _conf_info.ServiceRPC)
	}

	if _conf_info.LogPath = iniData.String("log_path"); _conf_info.LogPath == "" {
		_conf_info.LogPath = LOG_PATH

		tlog.Warn("reload srv config (%s) warn (default %s).", "log_path", _conf_info.LogPath)
	}

	if _conf_info.SentryDsn = iniData.String("sentry_dsn"); _conf_info.SentryDsn == "" {
		_conf_info.SentryDsn = SENTRY_DSN

		tlog.Warn("reload srv config (%s) warn (default %s).", "sentry_dsn", _conf_info.SentryDsn)
	}

	accntRegisterLimit, err := iniData.Int64("accnt_register_limit")

	if err != nil {
		_conf_info.AccntRegisterLimit = ACCNT_REGISTER_LIMIT

		tlog.Warn("reload srv config (%s) warn (default %d).", "accnt_register_limit", _conf_info.AccntRegisterLimit)
	} else {
		_conf_info.AccntRegisterLimit = uint64(accntRegisterLimit)
	}

	dbMaxIdleConn, err := iniData.Int("db_max_idle_conn")

	if err != nil {
		_conf_info.DBMaxIdleConn = DB_MAX_IDLE_CONN

		tlog.Warn("reload srv config (%s) warn (default %d).", "db_max_idle_conn", _conf_info.DBMaxIdleConn)
	} else {
		_conf_info.DBMaxIdleConn = dbMaxIdleConn
	}

	dbMaxOpenConn, err := iniData.Int("db_max_open_conn")

	if err != nil {
		_conf_info.DBMaxOpenConn = DB_MAX_OPEN_CONN

		tlog.Warn("reload srv config (%s) warn (default %d).", "db_max_open_conn", _conf_info.DBMaxOpenConn)
	} else {
		_conf_info.DBMaxOpenConn = dbMaxOpenConn
	}

	redisMaxIdleConn, err := iniData.Int("redis_max_idle_conn")

	if err != nil {
		_conf_info.RedisMaxIdleConn = REDIS_MAX_IDLE_CONN

		tlog.Warn("reload srv config (%s) warn (default %d).", "redis_max_idle_conn", _conf_info.RedisMaxIdleConn)
	} else {
		_conf_info.RedisMaxIdleConn = redisMaxIdleConn
	}

	redisTimeout, err := iniData.Int("redis_timeout")

	if err != nil {
		_conf_info.RedisTimeout = REDIS_TIMEOUT

		tlog.Warn("reload srv config (%s) warn (default %d).", "redis_timeout", _conf_info.RedisTimeout)
	} else {
		_conf_info.RedisTimeout = redisTimeout
	}

	memcachedMaxOpenConn, err := iniData.Int("memcached_max_open_conn")

	if err != nil {
		_conf_info.MemcachedMaxOpenConn = MEMCACHED_MAX_OPEN_CONN

		tlog.Warn("reload srv config (%s) warn (default %d).", "memcached_max_open_conn", _conf_info.MemcachedMaxOpenConn)
	} else {
		_conf_info.MemcachedMaxOpenConn = memcachedMaxOpenConn
	}

	if _conf_info.RunMode == "prod" {
		_conf_info.LoginDB = iniData.String("prod::login_db")
	} else {
		_conf_info.LoginDB = iniData.String("dev::login_db")
	}

	if _conf_info.LoginDB == "" {
		_conf_info.LoginDB = LOGIN_DB

		tlog.Warn("reload srv config (%s) warn (default %s).", "login_db", _conf_info.LoginDB)
	}

	if _conf_info.RunMode == "prod" {
		_conf_info.AccountRedisUrl = iniData.String("prod::account_redis_url")
	} else {
		_conf_info.AccountRedisUrl = iniData.String("dev::account_redis_url")
	}

	if _conf_info.AccountRedisUrl == "" {
		_conf_info.AccountRedisUrl = ACCOUNT_REDIS_URL

		tlog.Warn("reload srv config (%s) warn (default %s).", "account_redis_url", _conf_info.AccountRedisUrl)
	}

	if _conf_info.RunMode == "prod" {
		_conf_info.AccountRedisAuth = iniData.String("prod::account_redis_auth")
	} else {
		_conf_info.AccountRedisAuth = iniData.String("dev::account_redis_auth")
	}

	if _conf_info.AccountRedisAuth == "" {
		_conf_info.AccountRedisAuth = ACCOUNT_REDIS_AUTH

		tlog.Warn("reload srv config (%s) warn (default %s).", "account_redis_auth", _conf_info.AccountRedisAuth)
	}

	if _conf_info.RunMode == "prod" {
		_conf_info.MemcachedUrl = iniData.String("prod::memcached_url")
	} else {
		_conf_info.MemcachedUrl = iniData.String("dev::memcached_url")
	}

	if _conf_info.MemcachedUrl == "" {
		_conf_info.MemcachedUrl = MEMCACHED_URL

		tlog.Warn("reload srv config (%s) warn (default %s).", "memcached_url", _conf_info.MemcachedUrl)
	}
}
