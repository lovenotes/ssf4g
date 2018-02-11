package clientmodel

import (
	"net/http"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/common/resource-data"
)

// Func - 切换ResourceSrv
func ResourceSwitch(w http.ResponseWriter, channelid uint32, channelver string) *tlog.ErrData {
	resourceInfo, ret := resourcedata.GetResourceInfo(channelid, channelver)

	if ret == false {
		tlog.Error("resource switch model (%d, %s) err (resource info nil).", channelid, channelver)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "resource info nil")

		return nil
	}

	respData := map[string]interface{}{
		"resource_addr": resourceInfo.ResourceAddr,
		"comment":       resourceInfo.Comment,
	}

	respdata.BuildRespSuccessRetData(w, 0, respData)

	return nil
}
