package clienthandler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"ssf4g/common/tlog"
	"ssf4g/common/utility"
	"ssf4g/server/login-srv/handler/client-model"
)

// Func - 账号注册
func AccountRegister(wr http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	accntName := ""
	accntPass := ""

	strAccntNames := r.Form["accnt_name"]
	strAccntName := ""

	if strAccntNames != nil && len(strAccntNames) > 0 {
		strAccntName = strAccntNames[0]
	}

	if strAccntName != "" {
		accntName = strAccntName
	} else {
		tlog.Error("account register (%v) err (accnt_name nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "accnt_name illegal")

		return
	}

	strAccntPasses := r.Form["accnt_pass"]
	strAccntPass := ""

	if strAccntPasses != nil && len(strAccntPasses) > 0 {
		strAccntPass = strAccntPasses[0]
	}

	if strAccntPass != "" {
		accntPass = strAccntPass
	} else {
		tlog.Error("account register (%v) err (accnt_pass nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "accnt_pass illegal")

		return
	}

	realIP := utility.GetIP(r)

	// 注册渠道, tchannel
	account, errData := clientmodel.AccountRegister(accntName, accntPass, realIP)

	if errData != nil {

	}

}
