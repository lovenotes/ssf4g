package clientrouter

import (
	"ssf4g/server/resource-srv/handler/client-handler"

	"github.com/gorilla/mux"
)

func RegisterApiRouter(muxrouter *mux.Router) {
	muxrouter.HandleFunc("/resource/v1/switch", clienthandler.ResourceSwitch).Methods("POST")
	muxrouter.HandleFunc("/resource/v1/detail", clienthandler.ResourceDetail).Methods("POST")
	muxrouter.HandleFunc("/resource/v1/portals", clienthandler.ResourcePortals).Methods("POST")
}
