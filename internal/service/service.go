package service

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pract2/internal/repo"
)

type Task interface {
	CreateTask(ctx *fiber.Ctx) error
	GetAllTasks(ctx *fiber.Ctx) error
	GetTaskById(ctx *fiber.Ctx) error
	UpdateTaskById(ctx *fiber.Ctx) error
	DeleteTaskById(ctx *fiber.Ctx) error
}

type Service struct {
	Task Task
}

func NewService(repo *repo.Repository, logger *zap.SugaredLogger) *Service {
	return &Service{
		Task: NewTaskService(repo.Task, logger),
	}
}
