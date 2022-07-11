package composites

import (
	"ca-library-app/internal/adapters/api"
	author2 "ca-library-app/internal/adapters/api/author"
	author3 "ca-library-app/internal/adapters/db/author"
	"ca-library-app/internal/domain/author"
)

type AuthorComposite struct {
	Storage author.Storage
	Service author2.Service
	Handler api.Handler
}

func NewAuthorComposite(mongoComposite *MongoDBComposite) (*AuthorComposite, error) {
	storage := author3.NewAuthorStorage(mongoComposite.db)
	service := author.NewService(storage)
	handler := author2.NewHandler(service)
	return &AuthorComposite{
		Storage: storage,
		Service: service,
		Handler: handler,
	}, nil
}
