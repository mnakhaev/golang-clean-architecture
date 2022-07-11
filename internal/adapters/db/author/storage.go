package author

import (
	"ca-library-app/internal/domain/author"
	"ca-library-app/pkg/client/mongodb"
)

type authorStorage struct {
	db mongodb.Client
}

func NewAuthorStorage(db *mongodb.Client) author.Storage {
	return &authorStorage{db: *db}
}

func (a *authorStorage) GetOne(uuid string) *author.Author {
	return nil
}

func (a *authorStorage) GetAll(limit, offset int) []*author.Author {
	return nil
}

func (a *authorStorage) Create(author *author.Author) *author.Author {
	return nil
}

func (a *authorStorage) Delete(author *author.Author) error {
	return nil
}
