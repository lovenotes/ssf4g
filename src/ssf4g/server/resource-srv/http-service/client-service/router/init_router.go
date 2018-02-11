package clientrouter

import (
	"net/http"

	"ssf4g/server/resource-srv/http-service/client-service/controller"

	"github.com/gorilla/mux"
)

func InitClientRouter(muxrouter *mux.Router) {
	muxrouter.NotFoundHandler = http.HandlerFunc(clientcontroller.PageNotFound)
	muxrouter.MethodNotAllowedHandler = http.HandlerFunc(clientcontroller.PageNotFound)

	RegisterApiRouter(muxrouter)
}
