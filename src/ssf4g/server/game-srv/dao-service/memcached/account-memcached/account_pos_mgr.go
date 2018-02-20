package accountmemcached

import (
	"fmt"

	"ssf4g/common/com-const"
	"ssf4g/common/tlog"
)

// Func - 设置AccntPos - 注:此处缓存为7天
func (accountmemcached *AccountMemcached) SetAccntPos(accntid uint64, srvid uint32) *tlog.ErrData {
	memKey := fmt.Sprintf(ACCNT_MEMCACHED_KEY, ACCNT_POS_KEY, ACCNT_POS_VER, accntid)

	err := accountmemcached._memcached.Set(memKey, []byte(fmt.Sprintf("%d", srvid)), int32(7*comconst.SEC_PER_DAY))

	// MemCache异常
	if err != nil {
		errMsg := tlog.Error("set accnt pos (%d, %d) err (memcached set %v).", accntid, srvid, err)

		return tlog.NewErrData(err, errMsg)
	}

	return nil
}

// Func - 删除AccntPos
func (accountmemcached *AccountMemcached) DelAccntPos(accntid uint64) *tlog.ErrData {
	memKey := fmt.Sprintf(ACCNT_MEMCACHED_KEY, ACCNT_POS_KEY, ACCNT_POS_VER, accntid)

	err := accountmemcached._memcached.Delete(memKey)

	// MemCache异常
	if err != nil {
		errMsg := tlog.Error("del accnt pos (%d, %d) err (memcached del %v).", accntid, err)

		return tlog.NewErrData(err, errMsg)
	}

	return nil
}
