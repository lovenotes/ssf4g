package clientmodel

import (
	"net/http"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/resource-srv/common/resource-data"
)

// Func - 获取ResourceSrv信息
func ResourceDetail(w http.ResponseWriter, platformid uint32) *tlog.ErrData {
	loginInfo, ret := resourcedata.GetLoginInfo()

	if ret == false {
		tlog.Error("resource detail model (%d) err (login info nil).", platformid)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "login info nil")

		return nil
	}

	platformInfo, ret := resourcedata.GetPlatformInfo(platformid)

	if ret == false {
		tlog.Error("resource detail modle (%d) err (platform info nil).", platformid)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "platform info nil")

		return nil
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

	return nil
}
