package lckmatch

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gyu-young-park/lck_data_generator/repository"
)

const HTTP_ROUTER_SUFFIX_LCK_MATCH = "/lck-match"

type Router struct {
	handler *Handler
}

func NewRouter(repo repository.Repository) *Router {
	return &Router{handler: NewHandler(repo)}
}

func (r *Router) Route(mux *mux.Router) {
	subRouter := mux.PathPrefix(HTTP_ROUTER_SUFFIX_LCK_MATCH).Subrouter()
	subRouter.HandleFunc("/", r.handler.getAllMatch).Methods(http.MethodGet)
}
