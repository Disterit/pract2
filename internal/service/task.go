package service

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pract2/internal/dto"
	"pract2/internal/models"
	"pract2/internal/repo"
	"strconv"
)

type TaskService struct {
	repo   repo.Task
	logger *zap.SugaredLogger
}

func NewTaskService(repo repo.Task, logger *zap.SugaredLogger) *TaskService {
	return &TaskService{repo: repo, logger: logger}
}

func (s *TaskService) CreateTask(ctx *fiber.Ctx) error {
	var input models.Task

	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("error parsing body", "error", err)
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "Invalid request body")
	}

	taskID, err := s.repo.CreateTask(input)
	if err != nil {
		s.logger.Errorw("error creating task", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   map[string]int{"taskID": taskID},
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *TaskService) GetAllTasks(ctx *fiber.Ctx) error {

	tasks, err := s.repo.GetAllTasks()
	if err != nil {
		s.logger.Errorw("error getting all tasks", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   tasks,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *TaskService) GetTaskById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "invalid id")
	}

	task, err := s.repo.GetTaskById(id)
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   task,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *TaskService) UpdateTaskById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "invalid id")
	}

	var input models.Task
	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("error parsing body", "error", err)
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "invalid body")
	}

	err = s.repo.UpdateTaskById(id, input)
	if err != nil {
		s.logger.Errorw("error updating task", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *TaskService) DeleteTaskById(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "invalid id")
	}

	err = s.repo.DeleteTaskById(id)
	if err != nil {
		s.logger.Errorw("error deleting task", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
