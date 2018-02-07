package resourcedata

import (
	"strconv"

	"ssf4g/common/tlog"
	"ssf4g/gamedata/csv-data"
)

// 配置表名
const (
	RESOURCE_INFO_DATA = "resource_info_data"
)

const (
	RESOURCE_INFO_CHANNEL_ID  = "channel_id"
	RESOURCE_INFO_CHANNEL_VER = "channel_ver"
	RESOURCE_INFO_ADDR        = "resource_addr"
	RESOURCE_INFO_COMMENT     = "comment"
)

type ResourceInfoData struct {
	ChannelID    uint32
	ChannelVer   string
	ResourceAddr string
	Comment      string
}

type ResourceInfoDatas struct {
	_resource_info_detail map[uint32]map[string]*ResourceInfoData
}

var (
	_resource_info_datas *ResourceInfoDatas
)

func initResourceInfo() {
	_resource_info_datas = &ResourceInfoDatas{}

	reloadResourceInfo()
}

func reloadResourceInfo() {
	_resource_info_datas._resource_info_detail = make(map[uint32]map[string]*ResourceInfoData)

	tableInfo, ret := csvdata.GetTable(RESOURCE_INFO_DATA)

	if ret != csvdata.CSV_DATA_OK {
		tlog.Error("reload login info (%s) err (table not exist).", RESOURCE_INFO_DATA)

		return
	}

	for key, value := range tableInfo {
		resourceInfo := &ResourceInfoData{}

		// Channel ID
		channelID, err := strconv.ParseUint(value.Fields[RESOURCE_INFO_CHANNEL_ID], 10, 32)

		if err != nil {
			tlog.Error("reload resource info (%s, %s, %s) err (channel id parse %v).", RESOURCE_INFO_DATA, key, value.Fields[RESOURCE_INFO_CHANNEL_ID], err)

			continue
		}

		resourceInfo.ChannelID = uint32(channelID)

		// Channel Ver
		channelVer := value.Fields[RESOURCE_INFO_CHANNEL_VER]

		if channelVer == "" {
			tlog.Error("reload resource info (%s, %s) err (channle ver nil).", RESOURCE_INFO_DATA, key)

			continue
		}

		resourceInfo.ChannelVer = channelVer

		// Resource Addr
		resourceAddr := value.Fields[RESOURCE_INFO_ADDR]

		if resourceAddr == "" {
			tlog.Error("reload resource info (%s, %s) err (resource addr nil).", RESOURCE_INFO_DATA, key)

			continue
		}
		resourceInfo.ResourceAddr = resourceAddr

		// Comment
		comment := value.Fields[RESOURCE_INFO_COMMENT]

		if comment == "" {
			tlog.Error("reload resource info (%s, %s) err (comment nil).", RESOURCE_INFO_DATA, key)

			continue
		}

		resourceInfo.Comment = comment

		if _resource_info_datas._resource_info_detail[resourceInfo.ChannelID] == nil {
			_resource_info_datas._resource_info_detail[resourceInfo.ChannelID] = make(map[string]*ResourceInfoData)
		}

		_resource_info_datas._resource_info_detail[resourceInfo.ChannelID][resourceInfo.ChannelVer] = resourceInfo
	}
}

func GetResourceSvrInfo(channelid uint32, channelver string) (*ResourceInfoData, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	resourceInfos, ret := _resource_info_datas._resource_info_detail[channelid]

	if ret == false {
		return nil, false
	}

	resourceInfo, retData := resourceInfos[channelver]

	if retData == false {
		return nil, false
	}

	return resourceInfo, true
}
