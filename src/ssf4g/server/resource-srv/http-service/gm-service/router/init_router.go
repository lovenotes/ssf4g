package gmrouter

import (
	"net/http"

	"ssf4g/server/resource-srv/http-service/gm-service/controller"

	"github.com/gorilla/mux"
)

func InitGMRouter(muxrouter *mux.Router) {
	muxrouter.NotFoundHandler = http.HandlerFunc(gmcontroller.PageNotFound)
	muxrouter.MethodNotAllowedHandler = http.HandlerFunc(gmcontroller.PageNotFound)

	RegisterApiRouter(muxrouter)
}
