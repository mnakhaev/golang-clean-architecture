package book

import (
	"ca-library-app/internal/domain/book"
)

// bookStorage can contain DB client, wrappers and many other things.
type bookStorage struct {
	//client mongodb.Client
}

func NewBookStorage() book.Storage {
	return &bookStorage{}
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
