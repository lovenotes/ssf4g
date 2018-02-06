package clientrouter

import (
	"net/http"

	"ssf4g/server/resource-srv/handler/client-handler"

	"github.com/gorilla/mux"
)

func InitClientRouter(muxrouter *mux.Router) {
	muxrouter.NotFoundHandler = http.HandlerFunc(clienthandler.PageNotFound)
	muxrouter.MethodNotAllowedHandler = http.HandlerFunc(clienthandler.PageNotFound)

	RegisterApiRouter(muxrouter)
}
