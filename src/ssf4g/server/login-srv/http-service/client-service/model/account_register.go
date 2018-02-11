package clientmodel

import (
	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/common/err-code"
	"ssf4g/server/login-srv/database"
)

func AccountRegister(accntname, accntpass, realip string) (uint32, *tlog.ErrData) {
	accountDB, errData := dbmgr.GetLoginDao().FirstOrInitAccount()

	if err != nil {
		errMsg := tlog.Error("account register model (%s, %s, %s) err (database %v).", accntname, accntpass, realip, errData.Error())

		return nil, errData.AttachErrMsg(errMsg)
	}

	if accountDB != nil {
		tlog.Warn("account register model (%s, %s, %s) err (account exists).", accntname, accntpass, realip)

		return nil, errconst.LOGIN_ERR_TYPE_ACCOUNT_EXIST
	}

	accntInfo := &account.Account{}
	accntInfo.AccntName = accntname
	accntInfo.PassHash = crypto.Sha1Hash(accntname + pass)
	accntInfo.Email = email
	accntInfo.Phone = phone
	accntInfo.RealName = realname
	accntInfo.IDNumber = idnumber
	accntInfo.LastIP = curip
	accntInfo.Platform = platform
	accntInfo.CreateTime = time.Now()
	accntInfo.ModifyTime = time.Now()

	ticketID, ret := getUniqueAccntID()

	if ret != errcode.ERR_COM_SUCCESS {
		logger.GetNLog().Error("get unique accnt id (%s, %d, %s) err (%v).",
			accntname, platform, curip, ret)
		return nil, ret
	}

	logger.GetNLog().Debug("register accnt (%s, %d, %s) ticketid: %v success.", accntname, platform, curip, ticketID)

	_, err = _accnt_dao.CreateAccnt(accntInfo, ticketID)

	if err != nil {
		logger.GetNLog().Error("create accnt (%s, %d, %s) err (%v).", accntname, platform, curip, err)
		return nil, errcode.ERR_COM_DB_ERROR
	}
	accntInfo.AccntID = ticketID

	return accntInfo, errcode.ERR_COM_SUCCESS
}
