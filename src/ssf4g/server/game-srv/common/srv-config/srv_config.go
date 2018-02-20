package srvconfig

import (
	"sync"

	"ssf4g/common/config"
	"ssf4g/common/tlog"
)

const (
	SRV_NAME = "login_srv"

	ZONE_ID = 1

	SRV_ID  = 1
	SRV_VER = "0.0.1"

	RUN_MODE = "dev"

	SERVICE        = "0.0.0.0:8001"
	SERVICE_PORTAL = "0.0.0.0:8002"

	LOG_PATH   = "/data/ssf4g/logs/gamesrv.log"
	SENTRY_DSN = ""

	DB_MAX_IDLE_CONN = 10
	DB_MAX_OPEN_CONN = 100

	REDIS_MAX_IDLE_CONN = 10
	REDIS_TIMEOUT       = 5

	MEMCACHED_MAX_OPEN_CONN = 100

	GAME_DB = "ssf4g:ssf4g@(127.0.0.1:3306)/game?timeout=30s&parseTime=true&loc=Local&charset=utf8"

	SERVER_REDIS_URL  = "127.0.0.1:6379"
	SERVER_REDIS_AUTH = ""

	MEMCACHED_URL = "127.0.0.1:11211"
)

type SrvConfig struct {
	SrvName string
	ZoneID  uint32
	SrvID   uint32
	SrvVer  string
	RunMode string

	Service       string
	ServicePortal string

	LogPath   string
	SentryDsn string

	DBMaxIdleConn int
	DBMaxOpenConn int

	RedisMaxIdleConn int
	RedisTimeout     int

	MemcachedMaxOpenConn int

	GameDB string

	ServerRedisUrl  string
	ServerRedisAuth string

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

	zoneID, err := iniData.Int("zone_id")

	if err != nil {
		_conf_info.ZoneID = ZONE_ID

		tlog.Warn("reload srv config (%s) warn (default %d).", "zone_id", _conf_info.ZoneID)
	} else {
		_conf_info.ZoneID = uint32(zoneID)
	}

	srvID, err := iniData.Int("srv_id")

	if err != nil {
		_conf_info.SrvID = SRV_ID

		tlog.Warn("reload srv config (%s) warn (default %d).", "srv_id", _conf_info.SrvID)
	} else {
		_conf_info.SrvID = uint32(srvID)
	}

	if _conf_info.SrvVer = iniData.String("srv_ver"); _conf_info.SrvVer == "" {
		_conf_info.SrvVer = SRV_VER

		tlog.Warn("reload srv config (%s) warn (default %s).", "srv_ver", _conf_info.SrvVer)
	}

	if _conf_info.RunMode = iniData.String("run_mode"); _conf_info.RunMode == "" {
		_conf_info.RunMode = RUN_MODE

		tlog.Warn("reload srv config (%s) warn (default %s).", "run_mode", _conf_info.RunMode)
	}

	if _conf_info.Service = iniData.String("service"); _conf_info.Service == "" {
		_conf_info.Service = SERVICE

		tlog.Warn("reload srv config (%s) warn (default %s).", "service", _conf_info.Service)
	}

	if _conf_info.ServicePortal = iniData.String("service_portal"); _conf_info.ServicePortal == "" {
		_conf_info.ServicePortal = SERVICE_PORTAL

		tlog.Warn("reload srv config (%s) warn (default %s).", "service_portal", _conf_info.ServicePortal)
	}

	if _conf_info.LogPath = iniData.String("log_path"); _conf_info.LogPath == "" {
		_conf_info.LogPath = LOG_PATH

		tlog.Warn("reload srv config (%s) warn (default %s).", "log_path", _conf_info.LogPath)
	}

	if _conf_info.SentryDsn = iniData.String("sentry_dsn"); _conf_info.SentryDsn == "" {
		_conf_info.SentryDsn = SENTRY_DSN

		tlog.Warn("reload srv config (%s) warn (default %s).", "sentry_dsn", _conf_info.SentryDsn)
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
		_conf_info.GameDB = iniData.String("prod::game_db")
	} else {
		_conf_info.GameDB = iniData.String("dev::game_db")
	}

	if _conf_info.GameDB == "" {
		_conf_info.GameDB = GAME_DB

		tlog.Warn("reload srv config (%s) warn (default %s).", "game_db", _conf_info.GameDB)
	}

	if _conf_info.RunMode == "prod" {
		_conf_info.ServerRedisUrl = iniData.String("prod::server_redis_url")
	} else {
		_conf_info.ServerRedisUrl = iniData.String("dev::server_redis_url")
	}

	if _conf_info.ServerRedisUrl == "" {
		_conf_info.ServerRedisUrl = SERVER_REDIS_URL

		tlog.Warn("reload srv config (%s) warn (default %s).", "server_redis_url", _conf_info.ServerRedisUrl)
	}

	if _conf_info.RunMode == "prod" {
		_conf_info.ServerRedisAuth = iniData.String("prod::server_redis_auth")
	} else {
		_conf_info.ServerRedisAuth = iniData.String("dev::server_redis_auth")
	}

	if _conf_info.ServerRedisAuth == "" {
		_conf_info.ServerRedisAuth = SERVER_REDIS_AUTH

		tlog.Warn("reload srv config (%s) warn (default %s).", "server_redis_auth", _conf_info.ServerRedisAuth)
	}

	if _conf_info.MemcachedUrl == "" {
		_conf_info.MemcachedUrl = MEMCACHED_URL

		tlog.Warn("reload srv config (%s) warn (default %s).", "memcached_url", _conf_info.MemcachedUrl)
	}
}
