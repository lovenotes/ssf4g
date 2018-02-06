package resourcedata

import (
	"strconv"

	"ssf4g/common/tlog"
	"ssf4g/gamedata/csv-data"
)

// 配置表名
const (
	ZONE_INFO_DATA = "zone_info_data"
)

const (
	ZONE_INFO_ID     = "zone_id"
	ZONE_INFO_STATUS = "zone_status"
)

type ZoneInfoData struct {
	ZoneID     int32
	ZoneStatus int32
}

type ZoneInfoDatas struct {
	_zone_info_detail map[int32]*ZoneInfoData
}

var (
	_zone_info_datas *ZoneInfoDatas
)

func initZoneInfo() {
	_zone_info_datas = &ZoneInfoDatas{}

	reloadZoneInfo()
}

func reloadZoneInfo() {
	_zone_info_datas._zone_info_detail = make(map[int32]*ZoneInfoData)

	tableInfo, ret := gamedata.GetTable(ZONE_INFO_DATA)
	if ret != gamedata.GAME_DATA_OK {
		logger.GetNLog().Error("get table (%s) err.", ZONE_INFO_DATA)
		return
	}

	for key, value := range tableInfo {
		zoneInfo := &ZoneInfoData{}

		zone_id, err := strconv.ParseInt(key, 10, 32)
		if err != nil {
			logger.GetNLog().Error("get key (%s, %s) err (%v).", ZONE_INFO_DATA, key, err)

			continue
		}
		zoneInfo.ZoneID = int32(zone_id)

		// Zone状态
		zone_status, err := strconv.ParseInt(value.Fields[ZONE_INFO_STATUS], 10, 32)
		if err != nil {
			logger.GetNLog().Error("get zone_status (%s, %s, %s) parseint err (%v).", ZONE_INFO_DATA, key, value.Fields[ZONE_INFO_STATUS], err)

			continue
		}

		zoneInfo.ZoneStatus = int32(zone_status)

		_zone_info_datas._zone_info_detail[zoneInfo.ZoneID] = zoneInfo
	}
}

func GetZoneInfo(zoneid int32) (*ZoneInfoData, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	zoneInfo, ret := _zone_info_datas._zone_info_detail[zoneid]

	return zoneInfo, ret
}
