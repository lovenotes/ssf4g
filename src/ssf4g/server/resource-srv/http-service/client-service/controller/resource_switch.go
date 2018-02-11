package clientcontroller

import (
	"net/http"
	"strconv"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/http-service/client-service/model"
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
			tlog.Error("resource switch controller (%v) err (channel_id parse %v).", r, err)

			respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "channel_id illegal")

			return
		}

		channelID = uint32(srcChannelID)
	} else {
		tlog.Error("resource switch controller (%v) err (channel_id nil).", r)

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
		tlog.Error("resource switch controller (%v) err (channel_ver nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "channel_ver illegal")

		return
	}

	errData := clientmodel.ResourceSwitch(w, channelID, channelVer)

	if errData != nil {
		tlog.Error("resource switch controller (%v) err (model %v).", r, errData.Error())

		return
	}

	return
}
