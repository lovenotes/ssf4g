package accountmemcached

import (
	"fmt"
	"strconv"

	"ssf4g/common/tlog"
)

// Func - 获取AccntPos
func (accountmemcached *AccountMemcached) GetAccntPos(accntid uint64) (uint32, *tlog.ErrData) {
	memKey := fmt.Sprintf(ACCNT_MEMCACHED_KEY, ACCNT_POS_KEY, ACCNT_POS_VER, accntid)

	accntPosVal, err := accountmemcached._memcached.Get(memKey)

	// MemCache异常
	if err != nil {
		errMsg := tlog.Error("get accnt pos (%d) err (memcached get %v).", accntid, err)

		return 0, tlog.NewErrData(err, errMsg)
	}

	// 数据不存在,DB加载并写入Memcached
	if accntPosVal == nil {
		tlog.Debug("get accnt pos (%d) warn (accnt pos not exist).", accntid)

		return 0, nil
	}

	accntPos, err := strconv.ParseInt(string(accntPosVal), 10, 32)

	if err != nil {
		errMsg := tlog.Error("get accnt pos (%d) err (parse %v).", accntid, err)

		return 0, tlog.NewErrData(err, errMsg)
	}

	return uint32(accntPos), nil
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
