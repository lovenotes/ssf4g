package clientmodel

import (
	"net/http"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/portal-srv/common/err-code"
	"ssf4g/server/portal-srv/dao-service/memcached"
	//	"ssf4g/server/portal-srv/dao-service/redis"
)

// Func - 获取GameSrv详情
func GameSrvDetail(w http.ResponseWriter, accntid uint64, accnttoken string) *tlog.ErrData {
	// 校验AccntToken合法性
	accntToken, errData := memcachedmgr.GetAccountMemcached().GetAccntToken(accntid)

	if errData != nil {
		errMsg := tlog.Error("game srv detail model (%d) err (get accnt tokent %v).", accntid, errData.Error())

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_SERVER_ERROR, "account memcached err")

		return errData.AttachErrMsg(errMsg)
	}

	if accntToken != accnttoken {
		tlog.Warn("game srv detail model (%s, %s) warn (accnt token illegal).", accnttoken, accntToken)

		respData := map[string]interface{}{
			"err_code": errcode.SRV_DETAIL_ERR_CODE_TYPE_TOKEN_ILLEGAL,
		}

		respdata.BuildRespSuccessRetData(w, 0, respData)

		return nil
	}

	// 获取GameSrv信息
	return nil
}
