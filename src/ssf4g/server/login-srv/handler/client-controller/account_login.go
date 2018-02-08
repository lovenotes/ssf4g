package clienthandler

import (
	//	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"bosslove/common/consts"
	"bosslove/common/errcode"
	"bosslove/common/logger"
	"bosslove/common/utils"
	"bosslove/loginsvr/common/svrconfig"
	"bosslove/loginsvr/server/accntmgr"
	"bosslove/loginsvr/server/account"
	"bosslove/loginsvr/server/gamedatamgr"
)

func UserRegister(wr http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(wr, "Method Not Allowed", 405)
		return
	}

	res := map[string]interface{}{}
	defer httpRespWrite(r, wr, time.Now(), res)

	if svrconfig.GetConfig().OperatorType != consts.OPERATOR_TYPE_DEFAULT {
		logger.GetNLog().Error("req api illegal.")
		res["ret"] = errcode.ERR_COM_REQ_API_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_REQ_API_ILLEGAL)
		return
	}

	// param
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.GetNLog().Error("body readall err (%v).", err)
		res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
		return
	}

	pStr := string(body)

	params, err := url.ParseQuery(pStr)
	if err != nil {
		logger.GetNLog().Error("body (%s) parsequery err (%v).", string(body), err)
		res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
		return
	}

	/*
		login_type := params.Get("login_type")
		loginType, err := strconv.Atoi(login_type)
		if err != nil {
			logger.GetNLog().Error("get login_type (%s) atoi err (%v).", login_type, err)
			res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
			res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
			return
		}
	*/

	platform_type := params.Get("platform_type")
	platformType, err := strconv.Atoi(platform_type)
	if err != nil {
		logger.GetNLog().Error("get platform_type (%s) atoi err (%v).", platform_type, err)
		res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
		return
	}

	accntName := params.Get("accnt_name")
	if accntName == "" {
		logger.GetNLog().Error("get accnt_name err (data nil).")
		res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
		return
	}

	accntPass := params.Get("accnt_pass")
	if accntPass == "" {
		logger.GetNLog().Error("get accnt_pass err (data nil).")
		res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
		return
	}

	email := params.Get("email")
	phone := params.Get("phone")
	realName := params.Get("real_name")
	IDNumber := params.Get("id_number")

	telecomOper := params.Get("telecom_oper")

	regist_channel := params.Get("regist_channel")
	if regist_channel == "" {
		regist_channel = "0"
	}
	registChannel, err := strconv.Atoi(regist_channel)
	if err != nil {
		logger.GetNLog().Error("get regist_channel (%s) atoi err (%v).", regist_channel, err)
		res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
		res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
		return
	}

	currentIP := util.GetIP(r)

	/*
		if svrconfig.GetConfig().LoginType != consts.LOGIN_TYPE_DEFAULT &&
			svrconfig.GetConfig().LoginType != int32(loginType) {
			logger.GetNLog().Error("login_type (%d, %d) err (param illegal).", svrconfig.GetConfig().LoginType, loginType)
			res["ret"] = errcode.ERR_COM_PARAM_ILLEGAL
			res["msg"] = gamedatamgr.GetErrCodeMsg(errcode.ERR_COM_PARAM_ILLEGAL)
			return
		}
	*/

	// 注册渠道, tchannel
	account, ret := accntmgr.RegisterAccnt(accntName, accntPass, email, phone, realName, IDNumber, currentIP, uint8(platformType))

	res["ret"] = ret

	// 如果失败,附带错误消息
	if ret != errcode.ERR_COM_SUCCESS {
		res["msg"] = gamedatamgr.GetErrCodeMsg(ret)
		return
	}

	// TLOG 相关
	logger.GetTLog().AsyncSendTlog(logger.TLOG_TYPE_REGIST,
		logger.GetTLog().GenAccntRegist(accntName, account.AccntID, int32(platformType), telecomOper, int32(registChannel)),
		logger.GetTLog().MapIDToConIdx())
}
