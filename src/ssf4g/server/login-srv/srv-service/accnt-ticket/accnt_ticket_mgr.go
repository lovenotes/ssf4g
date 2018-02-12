package accntticketmgr

import (
	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/dao-service/database"
	"ssf4g/server/login-srv/dao-service/redis"

	"github.com/garyburd/redigo/redis"
)

// Func - 初始化AccntTicket信息
func InitAccntTicket() *tlog.ErrData {
	ticketID, errData := redismgr.GetAccountRedis().GetTicketID()

	if errData != nil {
		switch errData.Error() {
		case redis.ErrNil:
			break
		default:
			errMsg := tlog.Error("init accnt ticket err (get ticket id %v).", errData.Error())

			return errData.AttachErrMsg(errMsg)
		}
	}

	accntID, errData := dbmgr.GetLoginDao().GetMaxAccntID()

	if errData != nil {
		errMsg := tlog.Error("init accnt ticket err (get max accnt id %v).", errData.Error())

		return errData.AttachErrMsg(errMsg)
	}

	// 获取默认设置初始的TicketID
	initTicketID, errData := dbmgr.GetLoginDao().GetTicketID()

	if errData != nil {
		errMsg := tlog.Error("init accnt ticket err (get init tick id %v).", errData.Error())

		return errData.AttachErrMsg(errMsg)
	}

	// 重置初始的TicketID
	if accntID > ticketID {
		ticketID = accntID
	}

	if initTicketID > ticketID {
		ticketID = initTicketID

		tlog.Info("init accnt ticket info (%d).", ticketID)

		errData := redismgr.GetAccountRedis().SetTicketID(ticketID)

		if errData != nil {
			errMsg := tlog.Error("init accnt ticket err (get init tick id %v).", errData.Error())

			return errData.AttachErrMsg(errMsg)
		}
	}

	return nil
}
