package logindao

import (
	"ssf4g/common/tlog"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type LoginDao struct {
	_db *gorm.DB
}

// Func - 初始化AppDao
func (dao *LoginDao) InitLoginDao(databaseurl string, maxidleconn, maxopenconn int) *tlog.ErrData {
	var err error

	dao._db, err = gorm.Open("mysql", databaseurl)

	if err != nil {
		errMsg := tlog.Error("init login dao err (%v).", err)

		return tlog.NewErrData(err, errMsg)
	}

	//dao._db.LogMode(true)
	dao._db.DB().SetMaxIdleConns(maxidleconn)
	dao._db.DB().SetMaxOpenConns(maxopenconn)

	return nil
}
