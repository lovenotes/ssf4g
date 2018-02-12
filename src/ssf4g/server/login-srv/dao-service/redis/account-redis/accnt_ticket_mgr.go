package accountredis

import (
	"fmt"

	"ssf4g/common/tlog"
	"ssf4g/server/login-srv/common/srv-const"

	"github.com/garyburd/redigo/redis"
)

// Func - 设置AccntTicket值
func (dao *AccountRedis) SetTicketID(ticketid uint64) *tlog.ErrData {
	con := dao._pool.Get()
	defer con.Close()

	ticketKey := fmt.Sprintf(srvconst.AccntTicketKey)

	_, err := con.Do("SET", ticketKey, ticketid)

	if err != nil {
		errMsg := tlog.Error("set ticket id (%s, %d) err (redis set %v).", ticketKey, ticketid, err)

		return tlog.NewErrData(err, errMsg)
	}

	tlog.Info("set ticket id info (%s, %d).", ticketKey, ticketid)

	return nil
}

// Func - 获取当前AccntTicket值
func (dao *AccountRedis) GetTicketID() (uint64, *tlog.ErrData) {
	con := dao._pool.Get()
	defer con.Close()

	ticketKey := fmt.Sprintf(srvconst.AccntTicketKey)

	ticketID, err := redis.Uint64(con.Do("GET", ticketKey))

	if err != nil {
		errMsg := tlog.Error("get ticket id (%s) err (redis get %v).", ticketKey, err)

		return 0, tlog.NewErrData(err, errMsg)
	}

	tlog.Debug("get ticket id (%s) debug (%d).", ticketKey, ticketID)

	return ticketID, nil
}

// Func - 生成并获取新的TicketID
func (dao *AccountRedis) GenTicketID() (uint64, *tlog.ErrData) {
	con := dao._pool.Get()
	defer con.Close()

	ticketKey := fmt.Sprintf(srvconst.AccntTicketKey)

	ticketID, err := redis.Uint64(con.Do("INCR", ticketKey))

	if err != nil {
		errMsg := tlog.Error("gen ticket id (%s) err (redis incr %v).", ticketKey, err)

		return 0, tlog.NewErrData(err, errMsg)
	}

	tlog.Debug("gen ticket id (%s) warn (%d).", ticketKey, ticketID)

	return ticketID, nil
}
