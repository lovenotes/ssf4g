package resourcedata

import (
	"strconv"
)

import (
	"ssf4g/common/tlog"
	"ssf4g/gamedata/csv-data"
)

// 配置表名
const (
	PLATFORM_INFO_DATA = "platform_info_data"
)

const (
	PLATFORM_INFO_ID           = "platform_id"
	PLATFORM_INFO_DOWNLOAD_URL = "download_url"
	PLATFORM_INFO_VERSIONS     = "versions"
)

type PlatformInfoData struct {
	PlatformID  uint32
	DownloadUrl string
	Versions    string
}

type PlatformInfoDatas struct {
	_platform_info_detail map[uint32]*PlatformInfoData
}

var (
	_platform_info_datas *PlatformInfoDatas
)

func initPlatformInfo() {
	_platform_info_datas = &PlatformInfoDatas{}

	reloadPlatformInfo()
}

func reloadPlatformInfo() {
	_platform_info_datas._platform_info_detail = make(map[uint32]*PlatformInfoData)

	tableInfo, ret := csvdata.GetTable(PLATFORM_INFO_DATA)

	if ret != csvdata.CSV_DATA_OK {
		tlog.Error("reload platform info (%s) err (table not exist).", PLATFORM_INFO_DATA)

		return
	}

	for key, value := range tableInfo {
		platformInfo := &PlatformInfoData{}

		// Platform ID
		platformID, err := strconv.ParseUint(key, 10, 32)

		if err != nil {
			tlog.Error("reload platform info (%s, %s) err (key parse %v).", PLATFORM_INFO_DATA, key, err)

			continue
		}

		platformInfo.PlatformID = uint32(platformID)

		// Download Url
		platformInfo.DownloadUrl = value.Fields[PLATFORM_INFO_DOWNLOAD_URL]

		// Versions
		platformInfo.Versions = value.Fields[PLATFORM_INFO_VERSIONS]

		_platform_info_datas._platform_info_detail[platformInfo.PlatformID] = platformInfo
	}
}

func GetPlatformInfo(platformid uint32) (*PlatformInfoData, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	platformInfo, ret := _platform_info_datas._platform_info_detail[platformid]

	return platformInfo, ret
}
