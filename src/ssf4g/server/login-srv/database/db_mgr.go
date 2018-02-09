package dbmgr

import (
	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/common/srv-config"
	"ssf4g/server/login-srv/database/login-dao"
)

var (
	_login_dao *logindao.LoginDao
)

func init() {
	maxIdleConn := srvconfig.GetConfig().DBMaxIdleConn
	maxOpenConn := srvconfig.GetConfig().DBMaxOpenConn

	loginUrl := srvconfig.GetConfig().LoginDB

	_login_dao = &logindao.LoginDao{}

	errData := _login_dao.InitLoginDao(loginUrl, maxIdleConn, maxOpenConn)

	if errData != nil {
		errMsg := tlog.Error("init login dao (%s, %d, %d) err (%v).", loginUrl, maxIdleConn, maxOpenConn, errData.Error())

		tlog.AsyncSend(errData.AttachErrMsg(errMsg))

		return
	}

	return
}

func GetLoginDao() *logindao.LoginDao {
	return _login_dao
}
