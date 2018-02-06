package errhandler

import (
	"net/http"

	"taptap-kit/const/lib-const"
	"taptap-kit/data/resp-data"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(int(libconst.STATUS_CODE_TYPE_DESC[libconst.ERR_MSG_TYPE_NOT_FOUND]))

	respData := respdata.NewRespFailedRetDataFromData(libconst.LANG_TYPE_DEF, libconst.ERR_MSG_TYPE_NOT_FOUND, "", []string{"url"})

	w.Write(respData.ToJson())
}
