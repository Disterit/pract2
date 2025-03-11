package main

import (
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"log"
	"os"
	"os/signal"
	"pract2/internal/api"
	"pract2/internal/config"
	CustomLogger "pract2/internal/logger"
	"pract2/internal/repo"
	"pract2/internal/service"
	"sync"
)

func main() {
	if err := godotenv.Load(config.EnvPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Error processing config", zap.Error(err))
	}

	logger, err := CustomLogger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal("Error init logger", zap.Error(err))
	}

	mu := new(sync.Mutex)
	repos := repo.NewRepository(mu)
	services := service.NewService(repos, logger)
	app := api.NewRouters(&api.Routers{Service: services}, cfg.Rest.Token)

	go func() {
		logger.Infof("Starting server on %s", cfg.Rest.ListenAddress)
		if err := app.Listen(cfg.Rest.ListenAddress); err != nil {
			log.Fatal(errors.Wrap(err, "failed to start server"))
		}
	}()

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	logger.Info("Shutting down server...")
}
