package lckmatch

import (
	"net/http"

	"github.com/gyu-young-park/lck_data_generator/api/responser"
	"github.com/gyu-young-park/lck_data_generator/repository"
)

type Handler struct {
	repo repository.Repository
}

func NewHandler(repo repository.Repository) *Handler {
	return &Handler{repo: repo}
}

func (h *Handler) getAllMatch(res http.ResponseWriter, req *http.Request) {
	data, err := h.repo.Get(string(repository.ALL_MATCH))
	if err != nil {
		responser.Response(res, http.StatusInternalServerError, "Error in server, Can't get server data\n")
		return
	}
	responser.Response(res, http.StatusOK, data)
}
