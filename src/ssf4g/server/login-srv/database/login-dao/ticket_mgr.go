package logindao

import (
	"ssf4g/common/tlog"
)

type Ticket struct {
	TicketId uint64
}

func (dao *LoginDao) GetTicketID() (uint64, *tlog.ErrData) {
	tickets := make([]*Ticket, 0)

	retGorm := dao._db.Find(&tickets)

	if retGorm.Error != nil {
		errMsg := tlog.Error("get ticket id err (db %v).", retGorm.Error)

		return 0, tlog.NewErrData(retGorm.Error, errMsg)
	}

	if len(tickets) == 0 {
		return 0, nil
	}

	return tickets[0].TicketId, nil
}
