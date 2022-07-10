package author

import (
	"ca-library-app/internal/domain/author"
)

type authorStorage struct {
}

func NewAuthorStorage() author.Storage {
	return &authorStorage{}
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
