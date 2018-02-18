package accountmemcached

import (
	"fmt"

	"ssf4g/common/tlog"
)

// Func - 获取AccntToken
func (accountmemcached *AccountMemcached) GetAccntToken(accntid uint64) (string, *tlog.ErrData) {
	memKey := fmt.Sprintf(ACCNT_MEMCACHED_KEY, ACCNT_TOKEN_KEY, ACCNT_TOKEN_VER, accntid)

	accntTokenVal, err := accountmemcached._memcached.Get(memKey)

	// MemCache异常
	if err != nil {
		errMsg := tlog.Error("get accnt token (%d) err (memcached get %v).", accntid, err)

		return "", tlog.NewErrData(err, errMsg)
	}

	// 数据不存在,DB加载并写入Memcached
	if accntTokenVal == nil {
		tlog.Warn("get accnt token (%d) warn (accnt token not exist).", accntid)

		return "", nil
	}

	accntToken := string(accntTokenVal)

	return accntToken, nil
}
