package book

import (
	"ca-library-app/internal/domain/book"
	"ca-library-app/pkg/client/mongodb"
)

// bookStorage can contain DB client, wrappers and many other things.
type bookStorage struct {
	db mongodb.Client
}

func NewBookStorage(db *mongodb.Client) book.Storage {
	return &bookStorage{db: *db}
}

func (b *bookStorage) GetOne(uuid string) *book.Book {
	return nil
}

func (b *bookStorage) GetAll(limit, offset int) []*book.Book {
	return nil
}

func (b *bookStorage) Create(book *book.Book) *book.Book {
	return nil
}

func (b *bookStorage) Delete(book *book.Book) error {
	return nil
}
