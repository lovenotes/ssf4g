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
	ZoneID     uint32
	ZoneStatus int32
}

type ZoneInfoDatas struct {
	_zone_info_detail map[uint32]*ZoneInfoData
}

var (
	_zone_info_datas *ZoneInfoDatas
)

func initZoneInfo() {
	_zone_info_datas = &ZoneInfoDatas{}

	reloadZoneInfo()
}

func reloadZoneInfo() {
	_zone_info_datas._zone_info_detail = make(map[uint32]*ZoneInfoData)

	tableInfo, ret := csvdata.GetTable(ZONE_INFO_DATA)

	if ret != csvdata.CSV_DATA_OK {
		tlog.Error("reload zone info (%s) err (table not exist).", ZONE_INFO_DATA)

		return
	}

	for key, value := range tableInfo {
		zoneInfo := &ZoneInfoData{}

		//Zone ID
		zoneID, err := strconv.ParseUint(key, 10, 32)

		if err != nil {
			tlog.Error("reload zone info (%s, %s) err (key parse %v).", ZONE_INFO_DATA, key, err)

			continue
		}

		zoneInfo.ZoneID = uint32(zoneID)

		// Zone Status
		zoneStatus, err := strconv.ParseInt(value.Fields[ZONE_INFO_STATUS], 10, 32)

		if err != nil {
			tlog.Error("reload zone info (%s, %s, %s) err (zone status parse %v).", ZONE_INFO_DATA, key, value.Fields[ZONE_INFO_STATUS], err)

			continue
		}

		zoneInfo.ZoneStatus = int32(zoneStatus)

		_zone_info_datas._zone_info_detail[zoneInfo.ZoneID] = zoneInfo
	}
}

func GetZoneInfo(zoneid uint32) (*ZoneInfoData, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	zoneInfo, ret := _zone_info_datas._zone_info_detail[zoneid]

	return zoneInfo, ret
}
