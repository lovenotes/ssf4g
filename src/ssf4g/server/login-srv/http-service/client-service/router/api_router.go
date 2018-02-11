package clientrouter

import (
	"ssf4g/server/login-srv/handler/client-controller"

	"github.com/gorilla/mux"
)

func RegisterApiRouter(muxrouter *mux.Router) {
	muxrouter.HandleFunc("/account/v1/register", clientcontroller.AccountRegister).Methods("POST")
	muxrouter.HandleFunc("/account/v1/login", clientcontroller.AccountLogin).Methods("POST")
}
