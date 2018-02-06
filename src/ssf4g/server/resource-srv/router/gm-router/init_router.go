package gmrouter

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitClientRouter(muxrouter *mux.Router) {
	muxrouter.NotFoundHandler = http.HandlerFunc(errcontroller.PageNotFound)
	muxrouter.MethodNotAllowedHandler = http.HandlerFunc(errcontroller.PageNotFound)

	RegisterApiRouter(muxrouter)
}
