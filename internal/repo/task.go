package repo

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"pract2/internal/models"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

func (r *TaskRepository) CreateTask(task models.Task) (int, error) {

	return 0, nil
}

func (r *TaskRepository) GetAllTasks() (map[int]models.Task, error) {

	return map[int]models.Task{}, nil
}

func (r *TaskRepository) GetTaskById(id int) (models.Task, error) {

	return models.Task{}, nil
}

func (r *TaskRepository) UpdateTaskById(id int, task models.Task) error {

	return nil
}

func (r *TaskRepository) DeleteTaskById(id int) error {

	return nil
}
