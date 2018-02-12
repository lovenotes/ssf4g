package accntticket

import (
	"bosslove/common/consts"
	"bosslove/common/logger"
	"bosslove/common/rediss"
	"bosslove/common/utils"
	"bosslove/loginsvr/common/consts"
	"bosslove/loginsvr/common/svrconfig"
)

// Func - 初始化AccntTicket信息
func InitAccntTicket() (int64, error) {
	// get the ticket from redis
	con := rediss.GetPool().Get()
	defer con.Close()

	// 获取Redis内的最大AccntID
	ticketKey := util.GenRedisKeyInt32(consts.AccntGolbalTicket, zoneid)
	ticket, err := redis.Int64(con.Do("GET", ticketKey))
	switch {
	case err == redis.ErrNil:
		break
	case err != nil:
		logger.GetNLog().Error("get ticket id (%d, %s) redis err (%v).",
			zoneid, ticketKey, err)
		return 0, err
	}

	accntID := int64(-1)

	if svrconfig.GetConfig().OperatorType == consts.OPERATOR_TYPE_DEFAULT {
		// 获取DB内的最大AccntID
		accntID, err = _accnt_dao.GetMaxAccntID()
		if err != nil {
			logger.GetNLog().Error("get max accnt id (%d) db err (%v).", zoneid, err)
			return 0, err
		}
	} else if svrconfig.GetConfig().OperatorType == consts.OPERATOR_TYPE_TENCENT {
		// 获取YYBDB内的最大AccntID
		accntID, err = _accnt_dao.GetTencentMaxAccntID()
		if err != nil {
			logger.GetNLog().Error("get yyb max accnt id (%d) db err (%v).", zoneid, err)
			return 0, err
		}
	}

	// 获取默认设置初始的TicketID
	ticketID, err := _ticket_dao.GetTicketID()
	if err != nil {
		logger.GetNLog().Error("get ticket id (%d) db err (%v).", zoneid, err)
		return 0, err
	}

	// 重置初始的TicketID
	if accntID > ticketID {
		ticketID = accntID
	}

	if ticketID > ticket {
		ticket = ticketID
		logger.GetNLog().Notic("init ticket to redis (%v, %s)...", ticket, ticketKey)
		//_, err = redis.String(con.Do("SET", ticketKey, ticket))
		con.Do("SET", ticketKey, ticket)
		if con.Err() != nil {
			logger.GetNLog().Error("set ticket id (%v, %s) redis err (%v).", ticket, ticketKey, con.Err())
			return 0, err
		}
	}
	return ticket, nil
}
