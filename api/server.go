package api

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gyu-young-park/lck_data_generator/api/healthy"
	"github.com/gyu-young-park/lck_data_generator/api/lckmatch"
	"github.com/gyu-young-park/lck_data_generator/repository"
)

const HTTP_ROUTER_SUFFIX_V1 = "/api/v1"
const HTTP_SERVER_PORT = ":8888"

type Server struct {
	repo         repository.Repository
	router       *mux.Router
}

func NewHTTPServer() *Server {
	server := &Server{}
	server.router = mux.NewRouter()
	server.repo = repository.NewFileRepository()
	return server
}

func (server *Server) setUpRoute() {
	healthy.NewRouter().Route(server.router)
	lckmatch.NewRouter(server.repo).Route(server.router)
}

func (server *Server) StartServer() {
	fmt.Println("start server!")
	server.setUpRoute()
	http.ListenAndServe(HTTP_SERVER_PORT, server.router)
}
