package accountmemcached

import (
	"fmt"

	"ssf4g/common/com-const"
	"ssf4g/common/tlog"
)

// Func - 设置AccntToken - 注:此处缓存为7天
func (accountmemcached *AccountMemcached) SetAccntToken(accntid uint64, accnttoken string) *tlog.ErrData {
	memKey := fmt.Sprintf(ACCNT_MEMCACHED_KEY, ACCNT_TOKEN_KEY, ACCNT_TOKEN_VER, accntid)

	err := accountmemcached._memcached.Set(memKey, []byte(accnttoken), int32(7*comconst.SEC_PER_DAY))

	// MemCache异常
	if err != nil {
		errMsg := tlog.Error("set accnt token (%d, %s) err (memcached set %v).", accntid, accnttoken, err)

		return tlog.NewErrData(err, errMsg)
	}

	return nil
}
