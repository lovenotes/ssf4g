package clientmodel

import (
	"net/http"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/common/resource-data"
)

// Func - 获取ResourceSrv信息
func ResourcePortals(w http.ResponseWriter) *tlog.ErrData {
	portalInfos, ret := resourcedata.GetPortalInfos()

	if ret == false {
		tlog.Error("resource portals model err (resource info nil).")

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "resource info nil")

		return nil
	}

	respdata.BuildRespSuccessRetData(w, 0, portalInfos)

	return nil
}
