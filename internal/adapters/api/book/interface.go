package book

import (
	"context"

	"ca-library-app/internal/domain/book"
)

// Service interface is a `use case` in terms of clean architecture
type Service interface {
	GetByUUID(ctx context.Context, uuid string) *book.Book
	GetBooks(ctx context.Context, limit, offset int) []*book.Book
	Create(ctx context.Context, dto *book.CreateBookDTO) *book.Book
}
