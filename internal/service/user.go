package service

import (
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"pract2/internal/models"
	"pract2/internal/repo"
)

type UserService struct {
	repo   repo.User
	logger *zap.SugaredLogger
}

func NewUserService(repo repo.User, logger *zap.SugaredLogger) *UserService {
	return &UserService{repo: repo, logger: logger}
}

func (s *UserService) SingUp(ctx *fiber.Ctx) error {
	var input models.User

	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("Error parsing body", "error", err)
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	//err := s.repo.SingUp(input)
	//if err != nil {
	//	s.logger.Errorw("Error creating user", "error", err)
	//	return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	//}

	return nil
}

func (s *UserService) SingIn(ctx *fiber.Ctx) error {

	return nil
}
