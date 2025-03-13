package repo

import (
	"github.com/gofiber/fiber/v2"
	"pract2/models"
	"sync"
)

var memory = map[int]models.Task{}
var id int

type TaskRepository struct {
	mu *sync.Mutex
}

func NewTaskRepository(mu *sync.Mutex) *TaskRepository {
	return &TaskRepository{mu: mu}
}

func (r *TaskRepository) CreateTask(task models.Task) (int, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	id++
	memory[id] = task

	return id, nil
}

func (r *TaskRepository) GetAllTasks() (map[int]models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	return memory, nil
}

func (r *TaskRepository) GetTaskById(id int) (models.Task, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	task, ok := memory[id]
	if !ok {
		return models.Task{}, fiber.ErrNotFound
	}

	return task, nil
}

func (r *TaskRepository) UpdateTaskById(id int, task models.Task) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	memory[id] = task

	return nil
}

func (r *TaskRepository) DeleteTaskById(id int) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	delete(memory, id)

	return nil
}
