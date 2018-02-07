package clientcontroller

import (
	"net/http"
	"strconv"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/common/resource-data"
)

// Func - 切换ResourceSrv
func ResourceSwitch(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	channelID := uint32(0)
	channelVer := ""

	strChannelIDs := r.Form["channel_id"]
	strChannelID := ""

	if strChannelIDs != nil && len(strChannelIDs) > 0 {
		strChannelID = strChannelIDs[0]
	}

	if strChannelID != "" {
		srcChannelID, err := strconv.ParseUint(strChannelID, 10, 32)

		if err != nil {
			tlog.Error("resource switch (%v) err (channel_id parse %v).", r, err)

			respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "channel_id illegal")

			return
		}

		channelID = uint32(srcChannelID)
	} else {
		tlog.Error("resource switch (%v) err (channel_id nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "channel_id illegal")

		return
	}

	strChannelVers := r.Form["channel_ver"]
	strChannelVer := ""

	if strChannelVers != nil && len(strChannelVers) > 0 {
		strChannelVer = strChannelVers[0]
	}

	if strChannelVer != "" {
		channelVer = strChannelVer
	} else {
		tlog.Error("resource switch (%v) err (channel_ver nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "channel_ver illegal")

		return
	}

	tlog.Debug("resource switch (%d, %s).", channelID, channelVer)

	resourceInfo, ret := resourcedata.GetResourceInfo(channelID, channelVer)

	if ret == false {
		tlog.Error("resource switch (%v) err (resource info nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "resource info nil")

		return
	}

	respData := map[string]interface{}{
		"resource_addr": resourceInfo.ResourceAddr,
		"comment":       resourceInfo.Comment,
	}

	respdata.BuildRespSuccessRetData(w, 0, respData)

	return
}
