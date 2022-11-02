package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gyu-young-park/lck_data_generator/api/healthy"
	"github.com/gyu-young-park/lck_data_generator/repository"
)

const HTTP_ROUTER_SUFFIX_V1 = "/api/v1"

type Server struct {
	repo         repository.Repository
	router       *mux.Router
	healthRouter *healthy.Router
}

func NewHTTPServer() *Server {
	server := &Server{}
	server.router = mux.NewRouter()
	server.healthRouter = healthy.NewRouter()
	server.repo = repository.NewFileRepository(repository.DEFAULT_RECORDING_JSON_FILE_NAME)
	return server
}

func (server *Server) setUpRoute() {
	server.healthRouter.Route(server.router)
}

func (server *Server) StartServer() {
	fmt.Println("start server!")
	server.setUpRoute()
	http.ListenAndServe(":8080", server.router)
}
