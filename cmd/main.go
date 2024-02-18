package main

import (
	"ca-library-app/internal/composites"
	"ca-library-app/internal/config"
	"ca-library-app/pkg/logging"
	"context"
	"github.com/julienschmidt/httprouter"
)

func main() {
	//entry point
	logger := logging.GetLogger()

	logger.Info("config initializing")
	cfg := config.GetConfig()

	logger.Info("router initializing")
	router := httprouter.New()

	logger.Info("mongodb composite initializing")
	mongoDBC, err := composites.NewMongoDBComposite(context.Background(), cfg.MongoDB.Host, cfg.MongoDB.Port, cfg.MongoDB.Username, cfg.MongoDB.Password, cfg.MongoDB.Database, cfg.MongoDB.AuthDB)
	if err != nil {
		logger.Fatal("mongodb composite failed")
	}

	logger.Info("author composite initializing")
	authorComposite, err := composites.NewAuthorComposite(mongoDBC)
	if err != nil {
		logger.Fatal("author composite failed")
	}
	authorComposite.Handler.Register(router)

	logger.Info("book composite initializing")
	bookComposite, err := composites.NewBookComposite(mongoDBC)
	if err != nil {
		logger.Fatal("book composite failed")
	}
	bookComposite.Handler.Register(router)

	//start

}
