package clientcontroller

import (
	"net/http"
	"strconv"

	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/portal-srv/http-service/client-service/model"
)

func GameSrvDetail(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	accntID := uint64(0)
	accntToken := ""

	strAccntIDs := r.Form["accnt_id"]
	strAccntID := ""

	if strAccntIDs != nil && len(strAccntIDs) > 0 {
		strAccntID = strAccntIDs[0]
	}

	if strAccntID != "" {
		srcAccntID, err := strconv.ParseUint(strAccntID, 10, 64)

		if err != nil {
			tlog.Error("game srv detail controller (%v) err (accnt_id parse %v).", r, err)

			respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "accnt_id illegal")

			return
		}

		accntID = uint64(srcAccntID)
	} else {
		tlog.Error("game srv detail controller (%v) err (accnt_id nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "accnt_id illegal")

		return
	}

	strAccntTokens := r.Form["accnt_token"]
	strAccntToken := ""

	if strAccntTokens != nil && len(strAccntTokens) > 0 {
		strAccntToken = strAccntTokens[0]
	}

	if strAccntToken != "" {
		accntToken = strAccntToken
	} else {
		tlog.Error("game srv detail controller (%v) err (accnt_token nil).", r)

		respdata.BuildRespFailedRetData(w, httpconst.STATUS_CODE_TYPE_INVALID_REQ, "accnt_token illegal")

		return
	}

	errData := clientmodel.GameSrvDetail(w, accntID, accntToken)

	if errData != nil {
		tlog.Error("game srv detail controller (%v) err (model %v).", r, errData.Error())

		return
	}

	return
}
