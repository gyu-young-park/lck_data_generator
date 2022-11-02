package healthy

import (
	"net/http"

	"github.com/gorilla/mux"
)

const HTTP_ROUTER_SUFFIX_HEALTHY = "/health"

type Router struct {
	handler *Handler
}

func NewRouter() *Router {
	return &Router{handler: NewHandler()}
}

func (r *Router) Route(mux *mux.Router) {
	subRouter := mux.PathPrefix(HTTP_ROUTER_SUFFIX_HEALTHY).Subrouter()
	subRouter.HandleFunc("/", r.handler.health).Methods(http.MethodGet)
}
