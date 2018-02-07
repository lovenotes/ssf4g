package resourcedata

import (
	"sync"

	"ssf4g/gamedata/csv-data"
	"ssf4g/server/resource-srv/common/srv-config"
)

var (
	_lock sync.RWMutex
)

func InitResourceData() {
	_lock.Lock()
	defer _lock.Unlock()

	// 初始化CSV数据
	resoureDataPath := svrconfig.GetConfig().ResourceDataPath
	csvdata.InitCsvData(resoureDataPath)

	initResourceInfo()
	initLoginInfo()
	initPortalInfo()
	initPlatformInfo()
	initZoneInfo()
}

func ReloadResourceData() {
	_lock.Lock()
	defer _lock.Unlock()

	// 重新加载CSV数据
	resoureDataPath := svrconfig.GetConfig().ResourceDataPath
	csvdata.ReloadCsvData(resoureDataPath)

	reloadResourceInfo()
	reloadLoginInfo()
	reloadPortalInfo()
	reloadPlatformInfo()
	reloadZoneInfo()
}
