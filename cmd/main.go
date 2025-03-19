package main

import (
	"context"
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

	pool, err := repo.Connection(context.Background(), cfg.PostgresDB)
	if err != nil {
		log.Fatal("Error connect to database", zap.Error(err))
	}

	// прооверка соединения с бд
	if err := repo.CheckConnection(pool, logger); err != nil {
		log.Fatalf("Connection check failed: %v", err)
	}

	// создаём интерфейсы для репозитория
	taskRepository := repo.NewTaskRepository(pool)
	userRepository := repo.NewUserRepository(pool)

	// создаем интерфейс репозитория
	repos := repo.NewRepository(taskRepository, userRepository)

	// создаём интерфейсы для сервиса
	taskService := service.NewTaskService(repos.Task, logger)
	userService := service.NewUserService(repos.User, logger, cfg.Service)

	// создаем интерфейс сервиса
	services := service.NewService(taskService, userService)

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

	if err = app.Shutdown(); err != nil {
		logger.Errorw("error shutting down server", zap.Error(err))
	}

	if err = repo.CloseConnection(pool); err != nil {
		logger.Errorw("error closing connection", zap.Error(err))
	}

	logger.Info("Server stopped")
}
