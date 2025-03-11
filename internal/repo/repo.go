package repo

import (
	"pract2/models"
	"sync"
)

type Task interface {
	CreateTask(task models.Task) (int, error)
	GetAllTasks() (map[int]models.Task, error)
	GetTaskById(id int) (models.Task, error)
	UpdateTaskById(id int, task models.Task) error
	DeleteTaskById(id int) error
}

type Repository struct {
	Task Task
}

func NewRepository(mu *sync.Mutex) *Repository {
	return &Repository{
		Task: NewTaskRepository(mu),
	}
}
