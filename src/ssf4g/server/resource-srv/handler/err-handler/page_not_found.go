package errhandler

import (
	"net/http"

	"ssf4g/common/http-const"
	"ssf4g/gamedata/resp-data"
)

func PageNotFound(w http.ResponseWriter, r *http.Request) {
	respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_NOT_FOUND, "page not found")
}
