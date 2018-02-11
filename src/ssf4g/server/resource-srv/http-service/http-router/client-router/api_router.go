package clientrouter

import (
	"ssf4g/server/resource-srv/handler/client-controller"

	"github.com/gorilla/mux"
)

func RegisterApiRouter(muxrouter *mux.Router) {
	muxrouter.HandleFunc("/resource/v1/switch", clientcontroller.ResourceSwitch).Methods("POST")
	muxrouter.HandleFunc("/resource/v1/detail", clientcontroller.ResourceDetail).Methods("POST")
	muxrouter.HandleFunc("/resource/v1/portals", clientcontroller.ResourcePortals).Methods("POST")
}
