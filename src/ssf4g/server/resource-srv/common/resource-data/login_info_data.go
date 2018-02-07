package resourcedata

import (
	"strconv"

	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/gamedata/csv-data"
)

// 配置表名
const (
	LOGIN_INFO_DATA = "login_info_data"
)

const (
	LOGIN_INFO_ID            = "login_id"
	LOGIN_INFO_TYPE          = "login_type"
	LOGIN_INFO_ADDR          = "login_addr"
	LOGIN_INFO_CDN_ADDR      = "cdn_addr"
	LOGIN_INFO_COMP_VER_LOW  = "comp_ver_low"
	LOGIN_INFO_COMP_VER_HIGH = "comp_ver_high"
)

type LoginInfoData struct {
	LoginID     uint32
	LoginType   uint32
	LoginAddr   string
	CDNAddr     string
	CompVerLow  string
	CompVerHigh string
}

type LoginInfoDatas struct {
	_login_info_detail []*LoginInfoData
}

var (
	_login_info_datas *LoginInfoDatas
)

func initLoginInfo() {
	_login_info_datas = &LoginInfoDatas{}

	reloadLoginInfo()
}

func reloadLoginInfo() {
	_login_info_datas._login_info_detail = make([]*LoginInfoData, 0)

	tableInfo, ret := csvdata.GetTable(LOGIN_INFO_DATA)

	if ret != csvdata.CSV_DATA_OK {
		tlog.Error("reload login info (%s) err (table not exist).", LOGIN_INFO_DATA)

		return
	}

	for key, value := range tableInfo {
		loginInfo := &LoginInfoData{}

		// Login ID
		loginID, err := strconv.ParseUint(key, 10, 32)

		if err != nil {
			tlog.Error("reload login info (%s, %s) err (key parse %v).", LOGIN_INFO_DATA, key, err)

			continue
		}

		loginInfo.LoginID = uint32(loginID)

		// Login Type
		loginType, err := strconv.ParseUint(value.Fields[LOGIN_INFO_TYPE], 10, 32)

		if err != nil {
			tlog.Error("reload login info (%s, %s, %s) err (login type parse %v).", LOGIN_INFO_DATA, key, value.Fields[LOGIN_INFO_TYPE], err)

			continue
		}

		loginInfo.LoginType = uint32(loginType)

		// Login Addr
		loginAddr := value.Fields[LOGIN_INFO_ADDR]

		if loginAddr == "" {
			tlog.Error("reload login info (%s, %s) err (login addr nil).", LOGIN_INFO_DATA, key)

			continue
		}

		loginInfo.LoginAddr = loginAddr

		// CDN Addr
		cdnAddr := value.Fields[LOGIN_INFO_CDN_ADDR]

		if cdnAddr == "" {
			tlog.Error("reload login info (%s, %s) err (cdn addr nil).", LOGIN_INFO_DATA, key)

			continue
		}

		loginInfo.CDNAddr = cdnAddr

		// Comp Ver Low
		compVerLow := value.Fields[LOGIN_INFO_COMP_VER_LOW]

		if compVerLow == "" {
			tlog.Error("reload login info (%s, %s) err (comp ver low nil).", LOGIN_INFO_DATA, key)

			continue
		}

		loginInfo.CompVerLow = compVerLow

		// Comp Ver High
		compVerHigh := value.Fields[LOGIN_INFO_COMP_VER_HIGH]

		if compVerHigh == "" {
			tlog.Error("reload login info (%s, %s) err (comp ver high nil).", LOGIN_INFO_DATA, key)

			continue
		}

		loginInfo.CompVerHigh = compVerHigh

		_login_info_datas._login_info_detail = append(_login_info_datas._login_info_detail, loginInfo)
	}
}

func GetLoginInfo() (*LoginInfoData, bool) {
	_lock.RLock()
	defer _lock.RUnlock()

	loginInfoCnt := len(_login_info_datas._login_info_detail)

	if loginInfoCnt == 0 {
		return nil, false
	}

	if loginInfoCnt == 1 {
		return _login_info_datas._login_info_detail[0], true
	}

	loginIndex := utility.RandNum(loginInfoCnt)

	return _login_info_datas._login_info_detail[loginIndex], true
}
