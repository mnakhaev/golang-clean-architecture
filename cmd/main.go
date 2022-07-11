package main

import (
	"log"

	"ca-library-app/internal/composites"
	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

// entities - это модели
// use cases - сервисы / то, что умеет делать приложение (регистрация пользователя, ...)

func main() {
	logger := logrus.New()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoComposite, err := composites.NewMongoDBComposite()
	if err != nil {
		log.Fatal("failed to initialize mongo composite")
	}

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoComposite)
	if err != nil {
		log.Fatal("failed to initialize books")
	}
	bookComposite.Handler.Register(router)

	logger.Info("author composite initializing")
	authComposite, err := composites.NewAuthorComposite(mongoComposite)
	if err != nil {
		log.Fatal("failed to initialize authors")
	}
	authComposite.Handler.Register(router)

	// start app
}
