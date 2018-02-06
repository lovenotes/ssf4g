package respdata

import (
	"fmt"
	"net/http"

	"ssf4g/common/http-const"

	"github.com/json-iterator/go"
)

type RespSuccessRetData struct {
	RespRet   bool        `json:"success"`
	RespDatas interface{} `json:"data"`
}

func (respsuccessretdata *RespSuccessRetData) ToJson() []byte {
	respData, err := jsoniter.Marshal(respsuccessretdata)

	if err != nil {
		return nil
	}

	return respData
}

func BuildRespSuccessRetData(w http.ResponseWriter, cacheexpired int32, respdatas interface{}) {
	respData := &RespSuccessRetData{
		RespRet:   true,
		RespDatas: respdatas,
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	if cacheexpired > 0 {
		w.Header().Add("Cache-Control", fmt.Sprintf("max-age=%d, private", cacheexpired))
	}

	w.WriteHeader(int(httpconst.STATUS_CODE_TYPE_DESC[httpconst.STATUS_CODE_TYPE_OK]))

	w.Write(respData.ToJson())
}

type RespFailedData struct {
	ErrMsg  string `json:"error"`
	MsgDesc string `json:"msg"`
}

type RespFailedRetData struct {
	RespRet  bool            `json:"success"`
	RespData *RespFailedData `json:"data"`
}

func (respfailedretdata *RespFailedRetData) ToJson() []byte {
	respData, err := jsoniter.Marshal(respfailedretdata)

	if err != nil {
		return nil
	}

	return respData
}

func BuildRespBriefFailedRetData(w http.ResponseWriter, errmsg, desc string) {
	respData := &RespFailedRetData{
		RespRet: false,
		RespData: &RespFailedData{
			ErrMsg:  errmsg,
			MsgDesc: desc,
		},
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(int(httpconst.STATUS_CODE_TYPE_DESC[errmsg]))

	w.Write(respData.ToJson())
}
