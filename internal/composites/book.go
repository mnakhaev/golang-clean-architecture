package composites

import (
	"ca-library-app/internal/adapters/api"
	book2 "ca-library-app/internal/adapters/api/book"
	book3 "ca-library-app/internal/adapters/db/book"
	"ca-library-app/internal/domain/book"
)

type BookComposite struct {
	Storage book.Storage
	Service book2.Service
	Handler api.Handler
}

// NewBookComposite cannot accept DB as an argument - there will be dependency on realization.
// If not - dependency is only on interfaces.
func NewBookComposite(mongoComposite *MongoDBComposite) (*BookComposite, error) {
	storage := book3.NewBookStorage(mongoComposite.db)
	service := book.NewService(storage)
	handler := book2.NewHandler(service)
	return &BookComposite{Storage: storage, Service: service, Handler: handler}, nil
}
