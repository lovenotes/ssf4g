package gamedao

import (
	"ssf4g/common/tlog"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type GameDao struct {
	_db *gorm.DB
}

// Func - 初始化GameDao
func (dao *LoginDao) InitGameDao(databaseurl string, maxidleconn, maxopenconn int) *tlog.ErrData {
	var err error

	dao._db, err = gorm.Open("mysql", databaseurl)

	if err != nil {
		errMsg := tlog.Error("init game dao err (%v).", err)

		return tlog.NewErrData(err, errMsg)
	}

	//dao._db.LogMode(true)
	dao._db.DB().SetMaxIdleConns(maxidleconn)
	dao._db.DB().SetMaxOpenConns(maxopenconn)

	return nil
}
