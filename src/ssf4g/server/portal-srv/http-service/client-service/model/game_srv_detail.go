package clientmodel

import (
	"net/http"

	"ssf4g/common/crypto"
	"ssf4g/common/http-const"
	"ssf4g/common/tlog"
	"ssf4g/gamedata/resp-data"
	"ssf4g/server/portal-srv/common/err-code"
	"ssf4g/server/portal-srv/dao-service/database"
)

func AccountLogin(w http.ResponseWriter, accntname, accntpass, realip string) *tlog.ErrData {
}
