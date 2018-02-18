package clientmodel

import (
	"net/http"

	"ssf4g/common/crypto"
	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/login-srv/common/err-code"
	"ssf4g/server/login-srv/common/srv-config"
	"ssf4g/server/login-srv/dao-service/database"
	"ssf4g/server/login-srv/dao-service/redis"
)

func AccountRegister(w http.ResponseWriter, accntname, accntpass, realip string) *tlog.ErrData {
	accountDB, errData := dbmgr.GetLoginDao().FirstOrInitAccount(accntname)

	if errData != nil {
		errMsg := tlog.Error("account register model (%s) err (first init accnt %v).", accntname, errData.Error())

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_SERVER_ERROR, "account database err")

		return errData.AttachErrMsg(errMsg)
	}

	if accountDB != nil && accountDB.AccntId != 0 {
		tlog.Warn("account register model (%s) warn (accnt name exist).", accntname)

		respData := map[string]interface{}{
			"err_code": errcode.REGISTER_ERR_CODE_TYPE_ACCNT_EXIST,
		}

		respdata.BuildRespSuccessRetData(w, 0, respData)

		return nil
	}

	accountDB.PassHash = crypto.EncryptSha1Hash(accntname + accntpass)
	accountDB.LastIp = realip

	accntRegisterLimit := srvconfig.GetConfig().AccntRegisterLimit

	ticketID, errData := redismgr.GetAccountRedis().GetTicketID()

	if errData != nil {
		errMsg := tlog.Error("account register model err (get ticket id %v).", errData.Error())

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_SERVER_ERROR, "ticket redis err")

		return errData.AttachErrMsg(errMsg)
	}

	if accntRegisterLimit != 0 && ticketID >= accntRegisterLimit {
		tlog.Warn("account register model (%s) warn (register limit).", accntname)

		respData := map[string]interface{}{
			"err_code": errcode.REGISTER_ERR_CODE_TYPE_REGISTER_LIMIT,
		}

		respdata.BuildRespSuccessRetData(w, 0, respData)

		return nil
	}

	// real check the new user limits
	ticketID, errData = redismgr.GetAccountRedis().GenTicketID()

	if errData != nil {
		errMsg := tlog.Error("account register model err (gen ticket id %v).", errData.Error())

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_SERVER_ERROR, "ticket redis err")

		return errData.AttachErrMsg(errMsg)
	}

	accountDB.AccntId = ticketID

	errData = dbmgr.GetLoginDao().SaveAccount(accountDB)

	if errData != nil {
		errMsg := tlog.Error("account register model (%s, %d) err (save accnt %v).", accntname, ticketID, errData.Error())

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_SERVER_ERROR, "account database err")

		return errData.AttachErrMsg(errMsg)
	}

	respData := map[string]interface{}{
		"accnt_id": ticketID,
	}

	respdata.BuildRespSuccessRetData(w, 0, respData)

	return nil
}
