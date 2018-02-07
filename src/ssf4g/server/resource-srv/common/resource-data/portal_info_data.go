package resourcedata

import (
	"strconv"

	"ssf4g/common/tlog"
	"ssf4g/gamedata/csv-data"

	"github.com/json-iterator/go"
)

// 配置表名
const (
	PORTAL_INFO_DATA = "portal_info_data"
)

const (
	PORTAL_INFO_ID      = "portal_id"
	PORTAL_INFO_ADDR    = "portal_addr"
	PORTAL_INFO_COMMENT = "comment"
)

type PortalInfoData struct {
	PortalID   uint32 `json:"portal_id"`
	PortalAddr string `json:"portal_addr"`
	Comment    string `json:"comment"`
}

type PortalInfoDatas struct {
	_portal_info_detail map[uint32]*PortalInfoData
	_portal_info_list   []*PortalInfoData
}

var (
	_portal_info_datas *PortalInfoDatas
)

func initPortalInfo() {
	_portal_info_datas = &PortalInfoDatas{}

	reloadPortalInfo()
}

func reloadPortalInfo() {
	_portal_info_datas._portal_info_detail = make(map[uint32]*PortalInfoData)
	_portal_info_datas._portal_info_list = make([]*PortalInfoData, 0)

	tableInfo, ret := csvdata.GetTable(PORTAL_INFO_DATA)

	if ret != csvdata.CSV_DATA_OK {
		tlog.Error("reload portal info (%s) err (table not exist).", PORTAL_INFO_DATA)

		return
	}

	for key, value := range tableInfo {
		portalInfo := &PortalInfoData{}

		// Portal ID
		portalID, err := strconv.ParseUint(key, 10, 32)

		if err != nil {
			tlog.Error("reload portal info (%s, %s) err (key parse %v).", PORTAL_INFO_DATA, key, err)

			continue
		}

		portalInfo.PortalID = uint32(portalID)

		// Portal Addr
		portalAddr := value.Fields[PORTAL_INFO_ADDR]

		if portalAddr == "" {
			tlog.Error("reload portal info (%s, %s) err (portal addr nil).", PORTAL_INFO_DATA, key)

			continue
		}

		portalInfo.PortalAddr = portalAddr

		// Portal Comment
		comment := value.Fields[PORTAL_INFO_COMMENT]

		if comment == "" {
			tlog.Error("reload portal info (%s, %s) err (comment nil).", PORTAL_INFO_DATA, key)

			continue
		}

		portalInfo.Comment = comment

		_portal_info_datas._portal_info_detail[portalInfo.PortalID] = portalInfo
		_portal_info_datas._portal_info_list = append(_portal_info_datas._portal_info_list, portalInfo)
	}
}

func GetPortalInfo(portalid uint32) (*PortalInfoData, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	portalInfo, ret := _portal_info_datas._portal_info_detail[portalid]

	if ret == false {
		return nil, false
	}

	return portalInfo, true
}

func GetPortalInfos(channelid, logintype int32) (string, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	portalInfos, err := jsoniter.Marshal(_portal_info_datas._portal_info_list)

	if err != nil {
		tlog.Error("get portal infos err (marshal %v).", err)

		return "", false
	}

	return string(portalInfos), true
}
