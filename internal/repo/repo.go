package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"pract2/internal/models"
)

type User interface {
	SingUp(ctx context.Context, username, password string) error
	SingIn(ctx context.Context, username string) (models.User, error)
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

func NewRepository(pool *pgxpool.Pool) *Repository {
	return &Repository{
		Task: NewTaskRepository(pool),
		User: NewUserRepository(pool),
	}
}
