package clientrouter

import (
	"net/http"

	"ssf4g/server/resource-srv/handler/err-handler"

	"github.com/gorilla/mux"
)

func InitClientRouter(muxrouter *mux.Router) {
	muxrouter.NotFoundHandler = http.HandlerFunc(errhandler.PageNotFound)
	muxrouter.MethodNotAllowedHandler = http.HandlerFunc(errhandler.PageNotFound)

	RegisterApiRouter(muxrouter)
}
