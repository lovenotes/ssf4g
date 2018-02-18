package clientrouter

import (
	"ssf4g/server/portal-srv/http-service/client-service/controller"

	"github.com/gorilla/mux"
)

func RegisterApiRouter(muxrouter *mux.Router) {
	muxrouter.HandleFunc("/portal/v1/srv-detail", clientcontroller.GameSrvDetail).Methods("POST")
}
