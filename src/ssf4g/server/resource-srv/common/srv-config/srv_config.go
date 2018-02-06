package svrconfig

import (
	"sync"

	"ssf4g/common/config"
	"ssf4g/common/tlog"
)

const (
	RUN_MODE = "dev"

	SERVICE     = "0.0.0.0:8001"
	SERVICE_GM  = "0.0.0.0:8011"
	SERVICE_RPC = "0.0.0.0:8021"

	LOG_PATH = "/data/ssf4g/logs/resourcesvr.log"

	RESOURCE_DATA_PATH = "/data/ssf4g/data/resourcedata"

	SENTRY_DSN = ""
)

type SrvConfig struct {
	RunMode string

	Service    string
	ServiceGM  string
	ServiceRPC string

	LogPath string

	ResourceDataPath string

	SentryDsn string
}

var (
	_conf_info *SrvConfig
	_lock      sync.Mutex
)

func InitSrvConfig() {
	config.InitConfig()

	_conf_info = &SrvConfig{}

	ReloadSrvConfig()
}

func GetConfig() *SrvConfig {
	return _conf_info
}

func ReloadSrvConfig() {
	_lock.Lock()
	defer _lock.Unlock()

	config.ReloadConfig()

	iniData := config.GetIniData()

	if _conf_info.RunMode = iniData.String("run_mode"); _conf_info.RunMode == "" {
		_conf_info.RunMode = RUN_MODE

		tlog.Warn("reload srv config (%s) warn (default %s).", "run_mode", _conf_info.RunMode)
	}

	if _conf_info.Service = iniData.String("service"); _conf_info.Service == "" {
		_conf_info.Service = SERVICE

		tlog.Warn("reload srv config (%s) warn (default %s).", "service", _conf_info.Service)
	}

	if _conf_info.ServiceGM = iniData.String("service_gm"); _conf_info.ServiceGM == "" {
		_conf_info.ServiceGM = SERVICE_GM

		tlog.Warn("reload srv config (%s) warn (default %s).", "service_gm", _conf_info.ServiceGM)
	}

	if _conf_info.ServiceRPC = iniData.String("service_rpc"); _conf_info.ServiceRPC == "" {
		_conf_info.ServiceRPC = SERVICE_RPC

		tlog.Warn("reload srv config (%s) warn (default %s).", "service_rpc", _conf_info.ServiceRPC)
	}

	if _conf_info.LogPath = iniData.String("log_path"); _conf_info.LogPath == "" {
		_conf_info.LogPath = LOG_PATH

		tlog.Warn("reload srv config (%s) warn (default %s).", "log_path", _conf_info.LogPath)
	}

	if _conf_info.ResourceDataPath = iniData.String("resource_data_path"); _conf_info.ResourceDataPath == "" {
		_conf_info.ResourceDataPath = RESOURCE_DATA_PATH

		tlog.Warn("reload srv config (%s) warn (default %s).", "resource_data_path", _conf_info.ResourceDataPath)
	}

	if _conf_info.SentryDsn = iniData.String("sentry_dsn"); _conf_info.SentryDsn == "" {
		_conf_info.SentryDsn = SENTRY_DSN

		tlog.Warn("reload srv config (%s) warn (default %s).", "sentry_dsn", _conf_info.SentryDsn)
	}
}
