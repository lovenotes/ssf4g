package logindao

import (
	"time"

	"ssf4g/common/tlog"
)

type Account struct {
	AccntName string `gorm:"primary_key"`
	AccntId   uint64
	PassHash  string
	LastIp    string
	Platform  uint32
	CreatedAt time.Time
	UpdatedAt *time.Time
}

// Func - 获取AccntID最大值
func (dao *LoginDao) GetMaxAccntID() (uint64, *tlog.ErrData) {
	accounts := make([]*Account, 0)

	retGorm := dao._db.Order("`accnt_id` desc").Limit(1).Find(&accounts)

	if retGorm.Error != nil {
		errMsg := tlog.Error("get max accnt id err (db %v).", retGorm.Error)

		return 0, tlog.NewErrData(retGorm.Error, errMsg)
	}

	if len(accounts) == 0 {
		return 0, nil
	}

	return accounts[0].AccntId, nil
}

// Func - FirstOrInit账号信息
func (dao *LoginDao) FirstOrInitAccount(accntname string) (*Account, *tlog.ErrData) {
	account := &Account{}

	retGorm := dao._db.Where(Account{AccntName: accntname}).FirstOrInit(account)

	if retGorm.Error != nil {
		errMsg := tlog.Error("first or init account (%s) err (%v).", accntname, retGorm.Error)

		return nil, tlog.NewErrData(retGorm.Error, errMsg)
	}

	return account, nil
}

// Func - 保存账号信息
func (dao *LoginDao) SaveAccount(account *Account) *tlog.ErrData {
	retGorm := dao._db.Save(account)

	if retGorm.Error != nil {
		errMsg := tlog.Error("save account (%s) err (%v).", account.AccntName, retGorm.Error)

		return tlog.NewErrData(retGorm.Error, errMsg)
	}

	return nil
}

// Func - 根据name获取账号信息
func (dao *LoginDao) GetAccountByName(accntname string) (*Account, *tlog.ErrData) {
	accounts := make([]*Account, 0)

	retGorm := dao._db.Where("`accnt_name = ?`", accntname).Find(&accounts)

	if retGorm.Error != nil {
		errMsg := tlog.Error("get account by name (%s) err (db %v).", accntname, retGorm.Error)

		return nil, tlog.NewErrData(retGorm.Error, errMsg)
	}

	if len(accounts) == 0 {
		return nil, nil
	}

	return accounts[0], nil
}

// Func - 根据ID获取账号信息
func (dao *LoginDao) GetAccountByID(accntid uint64) (*Account, *tlog.ErrData) {
	accounts := make([]*Account, 0)

	retGorm := dao._db.Where("`accnt_id = ?`", accntid).Find(&accounts)

	if retGorm.Error != nil {
		errMsg := tlog.Error("get account by id (%d) err (db %v).", accntid, retGorm.Error)

		return nil, tlog.NewErrData(retGorm.Error, errMsg)
	}

	if len(accounts) == 0 {
		return nil, nil
	}

	return accounts[0], nil
}
