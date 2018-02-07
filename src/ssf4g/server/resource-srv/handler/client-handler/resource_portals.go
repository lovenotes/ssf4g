package clienthandler

import (
	"net/http"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/common/resource-data"
)

// Func - 获取PortalSrv列表
func ResourcePortals(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	portalInfos, ret := resourcedata.GetPortalInfos()

	if ret == false {
		tlog.Error("resource switch (%v) err (resource info nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "resource info nil")

		return
	}

	respdata.BuildRespSuccessRetData(w, 0, portalInfos)

	return
}
