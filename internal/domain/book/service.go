package book

import (
	"context"

	"ca-library-app/internal/adapters/api/author"
	"ca-library-app/internal/adapters/api/book"
)

// service struct should contain the main - it is storage (gateway).
type service struct {
	storage       Storage
	authorService author.Service // connecting between book service and author service
}

func NewService(storage Storage) book.Service {
	return &service{storage: storage}
}

func (s *service) GetByUUID(ctx context.Context, uuid string) *Book {
	return s.storage.GetOne(uuid)
}

func (s *service) GetBooks(ctx context.Context, limit, offset int) []*Book {
	return s.storage.GetAll(limit, offset)
}

func (s *service) Create(ctx context.Context, dto *CreateBookDTO) *Book {
	//author, err := s.authorService.GetByUUID(ctx, dto.AuthorUUID)
	// TODO return err if author doesn't exist
	return nil
}
