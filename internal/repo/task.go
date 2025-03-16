package repo

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"pract2/internal/models"
)

const (
	CreateTaskQuery  = `INSERT INTO tasks (user_id, title, description) VALUES ($1, $2, $3) RETURNING id`
	GetAllTasksQuery = `SELECT t.* FROM users AS u LEFT JOIN tasks AS t ON t.user_id = u.id WHERE u.username = $1`
	GetTaskByIdQuery = `SELECT * FROM tasks WHERE id = $1`
	UpdateTaskById   = `UPDATE tasks SET status = $1 WHERE id = $2 AND user_id = $3`
	DeleteTaskById   = `DELETE FROM tasks WHERE id = $1 AND user_id = $2`
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{pool: pool}
}

func (r *TaskRepository) CreateTask(ctx context.Context, task models.Task) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx, CreateTaskQuery, task.UserId, task.Title, task.Description).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *TaskRepository) GetAllTasks(ctx context.Context, username string) ([]models.Task, error) {

	var tasks []models.Task

	rows, err := r.pool.Query(ctx, GetAllTasksQuery, username)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var task models.Task
		err = rows.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (r *TaskRepository) GetTaskById(ctx context.Context, taskId int) (models.Task, error) {
	row := r.pool.QueryRow(ctx, GetTaskByIdQuery, taskId)
	var task models.Task
	err := row.Scan(&task.Id, &task.UserId, &task.Title, &task.Description, &task.Status, &task.CreatedAt)
	if err != nil {
		return models.Task{}, err
	}
	return task, nil
}

func (r *TaskRepository) UpdateTaskById(ctx context.Context, status string, taskId, userId int) error {
	_, err := r.pool.Exec(ctx, UpdateTaskById, status, taskId, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *TaskRepository) DeleteTaskById(ctx context.Context, taskId, userId int) error {
	_, err := r.pool.Exec(ctx, DeleteTaskById, taskId, userId)
	if err != nil {
		return err
	}
	return nil
}
