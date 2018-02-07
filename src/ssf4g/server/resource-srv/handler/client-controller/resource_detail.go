package clientcontroller

import (
	"net/http"
	"strconv"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/common/resource-data"
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
			tlog.Error("resource detail (%v) err (platform_id parse %v).", r, err)

			respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "platform_id illegal")

			return
		}

		platformID = uint32(srcPlatformID)
	} else {
		tlog.Error("resource detail (%v) err (platform_id nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "platform_id illegal")

		return
	}

	tlog.Debug("resource detail (%d).", platformID)

	loginInfo, ret := resourcedata.GetLoginInfo()

	if ret == false {
		tlog.Error("resource detail (%v) err (login info nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "login info nil")

		return
	}

	platformInfo, ret := resourcedata.GetPlatformInfo(platformID)

	if ret == false {
		tlog.Error("resource detail (%v) err (platform info nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "platform info nil")

		return
	}

	respData := map[string]interface{}{
		"login_type":    loginInfo.LoginType,
		"cdn_addr":      loginInfo.CDNAddr,
		"comp_ver_low":  loginInfo.CompVerLow,
		"comp_ver_high": loginInfo.CompVerHigh,
		"download_url":  platformInfo.DownloadUrl,
		"versions":      platformInfo.Versions,
	}

	respdata.BuildRespSuccessRetData(w, 0, respData)

	return
}
