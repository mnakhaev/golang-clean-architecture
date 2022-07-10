package book

import (
	"net/http"

	"ca-library-app/internal/adapters/api"
	"github.com/julienschmidt/httprouter"
)

const (
	bookURL  = "/books/:book_id"
	booksURL = "/books"
)

type handler struct {
	bookService Service
}

func NewHandler(service Service) api.Handler {
	return &handler{bookService: service}
}

func (h *handler) Register(router *httprouter.Router) {
	router.GET(bookURL, h.GetAllBooks)
}

func (h *handler) GetAllBooks(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	//books := h.bookService.GetBooks(context.Background(), 0, 0)
	// TODO: unmarshal books in bytes
	_, _ = w.Write([]byte("books"))
	w.WriteHeader(http.StatusOK)

}
