package clientcontroller

import (
	"net/http"
	"strconv"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/http-service/client-service/model"
)

// Func - 获取ResourceSrv信息
func ResourceDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	platformID := uint32(0)

	strPlatformIDs := r.Form["platform_id"]
	strPlatformID := ""

	if strPlatformIDs != nil && len(strPlatformIDs) > 0 {
		strPlatformID = strPlatformIDs[0]
	}

	if strPlatformID != "" {
		srcPlatformID, err := strconv.ParseUint(strPlatformID, 10, 32)

		if err != nil {
			tlog.Error("resource detail controller (%v) err (platform_id parse %v).", r, err)

			respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "platform_id illegal")

			return
		}

		platformID = uint32(srcPlatformID)
	} else {
		tlog.Error("resource detail controller (%v) err (platform_id nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "platform_id illegal")

		return
	}

	errData := clientmodel.ResourceDetail(w, platformID)

	if errData != nil {
		tlog.Error("resource detail controller (%v) err (model %v).", r, errData.Error())

		return
	}

	return
}
