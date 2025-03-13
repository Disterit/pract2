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
	// подгрузка env файла
	if err := godotenv.Load(config.EnvPath); err != nil {
		log.Fatal("Error loading .env file")
	}

	// помещаем env файл в структуру
	var cfg config.AppConfig
	if err := envconfig.Process("", &cfg); err != nil {
		log.Fatal("Error processing config", zap.Error(err))
	}

	// добавляем свой логгер
	logger, err := CustomLogger.NewLogger(cfg.LogLevel)
	if err != nil {
		log.Fatal("Error init logger", zap.Error(err))
	}

	// мьютекс для того чтобы данные были достоверные
	mu := new(sync.Mutex)
	// создаем стуктуру репозитория
	repos := repo.NewRepository(mu)
	// создаем сервиса репозитория
	services := service.NewService(repos, logger)
	// создаем хендлера репозитория
	app := api.NewRouters(&api.Routers{Service: services}, cfg.Rest.Token)

	//запускаем апи в горутине, чтобы при остановке отработали все задачи в очереди
	go func() {
		logger.Infof("Starting server on %s", cfg.Rest.ListenAddress)
		if err := app.Listen(cfg.Rest.ListenAddress); err != nil {
			log.Fatal(errors.Wrap(err, "failed to start server"))
		}
	}()

	// ждем завершение процесса
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt)
	<-signalChan

	logger.Info("Shutting down server...")
}
