package main

import (
	author3 "ca-library-app/internal/adapters/api/author"
	book3 "ca-library-app/internal/adapters/api/book"
	author2 "ca-library-app/internal/adapters/db/author"
	book2 "ca-library-app/internal/adapters/db/book"
	"ca-library-app/internal/domain/author"
	"ca-library-app/internal/domain/book"
)

// entities - это модели
// use cases - сервисы / то, что умеет делать приложение (регистрация пользователя, ...)

func main() {
	bookStorage := book2.NewBookStorage()
	bookService := book.NewService(bookStorage)
	bookHandler := book3.NewHandler(bookService)

	authorStorage := author2.NewAuthorStorage()
	authorService := author.NewService(authorStorage)
	authorHandler := author3.NewHandler(authorService)

	// register handlers in router
}
