package clientcontroller

import (
	"net/http"

	"ssf4g/common/tlog"
	"ssf4g/server/resource-srv/http-service/client-service/model"
)

// Func - 获取PortalSrv列表
func ResourcePortals(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	errData := clientmodel.ResourcePortals(w)

	if errData != nil {
		tlog.Error("resource portals controller (%v) err (model %v).", r, errData.Error())

		return
	}

	return
}
