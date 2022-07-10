package author

import (
	"net/http"

	"ca-library-app/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
)

const (
	authorURL  = "/authors/:author_id"
	authorsURL = "/authors"
)

type handler struct {
	authorService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{authorService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(authorURL, h.GetAllAuthors)
}

func (h *handler) GetAllAuthors(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//authors := h.authorService.GetBooks(context.Background(), 0, 0)
	// TODO: unmarshal books in bytes
	_, _ = w.Write([]byte("authors"))
	w.WriteHeader(http.StatusOK)

}
