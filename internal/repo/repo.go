package repo

import (
	"github.com/gofiber/fiber/v2"
	"pract2/internal/models"
	"sync"
)

type User interface {
	SingUp(ctx *fiber.Ctx) error
	SingIn(ctx *fiber.Ctx) error
}

type Task interface {
	CreateTask(task models.Task) (int, error)
	GetAllTasks() (map[int]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	UpdateTaskById(id int, task models.Task) error
	DeleteTaskById(id int) error
}

type Repository struct {
	Task Task
	User User
}

func NewRepository(mu *sync.Mutex) *Repository {
	return &Repository{
		Task: NewTaskRepository(mu),
	}
}
