package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"pract2/internal/models"
)

type User interface {
	SingUp(ctx context.Context, username, password string) error
	SingIn(ctx context.Context, username string) (models.User, error)
	DeleteUser(ctx context.Context, id int) error
}

type Task interface {
	CreateTask(ctx context.Context, task models.Task) (int, error)
	GetAllTasks(ctx context.Context, username string) ([]models.Task, error)
	GetTaskById(ctx context.Context, taskId int) (models.Task, error)
	UpdateTaskById(ctx context.Context, status string, taskId, userId int) error
	DeleteTaskById(ctx context.Context, taskId, userId int) error
}

type IRepository interface {
	GetTaskRepo() Task
	GetUserRepo() User
}

type Repository struct {
	Task Task
	User User
}

func (r *Repository) GetTaskRepo() Task {
	return r.Task
}

func (r *Repository) GetUserRepo() User {
	return r.User
}

func NewRepository(pool *pgxpool.Pool) IRepository {
	return &Repository{
		Task: NewTaskRepository(pool),
		User: NewUserRepository(pool),
	}
}
