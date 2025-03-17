package service

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pract2/internal/config"
	"pract2/internal/repo"
)

type User interface {
	SingUp(ctx *fiber.Ctx) error
	SingIn(ctx *fiber.Ctx) error
	DeleteUser(ctx *fiber.Ctx) error
}

type Task interface {
	CreateTask(ctx *fiber.Ctx) error
	GetAllTasks(ctx *fiber.Ctx) error
	GetTaskById(ctx *fiber.Ctx) error
	UpdateTaskById(ctx *fiber.Ctx) error
	DeleteTaskById(ctx *fiber.Ctx) error
}

type IService interface {
	GetTaskService() Task
	GetUserService() User
}

type Service struct {
	Task Task
	User User
}

func (s *Service) GetTaskService() Task {
	return s.Task
}

func (s *Service) GetUserService() User {
	return s.User
}

func NewService(repo repo.IRepository, logger *zap.SugaredLogger, cfg config.Service) IService {
	return &Service{
		Task: NewTaskService(repo.GetTaskRepo(), logger),
		User: NewUserService(repo.GetUserRepo(), logger, cfg),
	}
}
