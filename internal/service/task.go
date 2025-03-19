package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"pract2/internal/dto"
	"pract2/internal/models"
	"pract2/internal/repo"
	"strconv"
)

type taskService struct {
	repo   repo.Task
	logger *zap.SugaredLogger
}

func NewTaskService(repo repo.Task, logger *zap.SugaredLogger) Task {
	return &taskService{
		repo:   repo,
		logger: logger,
	}
}

func (s *taskService) CreateTask(ctx *fiber.Ctx) error {
	var input models.Task

	userId := ctx.Locals("user_id").(int)
	input.UserId = userId

	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("error parsing body", "error", err)
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "error")
	}

	taskID, err := s.repo.CreateTask(ctx.Context(), input)
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

func (s *taskService) GetAllTasks(ctx *fiber.Ctx) error {

	username := ctx.Locals("username").(string)

	tasks, err := s.repo.GetAllTasks(ctx.Context(), username)
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

// тут я просто написал по тз, но вообще не нужно делать проверку ради проверки объясню в readme
func (s *taskService) GetTaskById(ctx *fiber.Ctx) error {

	userId := ctx.Locals("user_id").(int)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "error")
	}

	task, err := s.repo.GetTaskById(ctx.Context(), id)
	if err != nil {
		if err == pgx.ErrNoRows {
			s.logger.Errorw("error getting task by id", "error", err)
			return dto.BadResponseError(ctx, dto.FieldBadFormat, "error")
		}
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.InternalServerError(ctx)
	}

	if task.UserId != userId {
		s.logger.Errorw("error no rights", "error", err)
		return dto.NotFoundError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   task,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *taskService) UpdateTaskById(ctx *fiber.Ctx) error {

	userId := ctx.Locals("user_id").(int)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "error")
	}

	var input models.Task
	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("error parsing body", "error", err)
		return dto.BadResponseError(ctx, dto.FieldIncorrect, "error")
	}

	err = s.repo.UpdateTaskById(ctx.Context(), input.Status, id, userId)
	if err != nil {
		s.logger.Errorw("error updating task", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (s *taskService) DeleteTaskById(ctx *fiber.Ctx) error {

	userId := ctx.Locals("user_id").(int)

	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("error getting task by id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, "error")
	}

	err = s.repo.DeleteTaskById(ctx.Context(), id, userId)
	if err != nil {
		s.logger.Errorw("error deleting task", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
