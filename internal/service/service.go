package service

import (
	"github.com/gofiber/fiber/v2"
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

type Service struct {
	Task Task
	User User
}

func NewService(task Task, user User) *Service {
	return &Service{
		Task: task,
		User: user,
	}
}
