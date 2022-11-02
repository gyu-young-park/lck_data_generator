package healthy

import (
	"net/http"

	"github.com/gyu-young-park/lck_data_generator/api/responser"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) health(res http.ResponseWriter, req *http.Request) {
	responser.Response(res, http.StatusOK, "Healthy check success!\n")
}
