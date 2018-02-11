package gmrouter

import (
	"net/http"

	"ssf4g/server/resource-srv/handler/err-controller"

	"github.com/gorilla/mux"
)

func InitGMRouter(muxrouter *mux.Router) {
	muxrouter.NotFoundHandler = http.HandlerFunc(errcontroller.PageNotFound)
	muxrouter.MethodNotAllowedHandler = http.HandlerFunc(errcontroller.PageNotFound)

	RegisterApiRouter(muxrouter)
}
