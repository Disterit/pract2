package service

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
	"pract2/internal/config"
	"pract2/internal/dto"
	"pract2/internal/models"
	"pract2/internal/repo"
	"strconv"
	"time"
)

type UserService struct {
	repo         repo.User
	logger       *zap.SugaredLogger
	passwordSalt string
	JWTSecret    string
}

func NewUserService(repo repo.User, logger *zap.SugaredLogger, cfg config.Service) *UserService {
	return &UserService{
		repo:         repo,
		logger:       logger,
		passwordSalt: cfg.PasswordSalt,
		JWTSecret:    cfg.Token,
	}
}

// регистрация пользователя
func (s *UserService) SingUp(ctx *fiber.Ctx) error {
	var input models.User

	// парсим данные пользователя
	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("Error parsing body", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, err.Error())
	}

	// хэшируем пароль для БД
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Errorw("Error hashing password", "error", err)
		return dto.InternalServerError(ctx)
	}

	// передаем данные имя и новый хешированный пароль
	err = s.repo.SingUp(ctx.Context(), input.Username, string(hashPassword))
	if err != nil {
		s.logger.Errorw("Error creating user", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// структура написана для задания, потому что в одном из заданий был пункт что нужно найти все задания пользователя по username
type CustomClaims struct {
	jwt.RegisteredClaims
	UserId   int    `json:"user_id"`
	Username string `json:"username"`
}

// функция для входа(получения jwt токена)
func (s *UserService) SingIn(ctx *fiber.Ctx) error {
	var input models.User

	// парсим данные для входа
	if err := ctx.BodyParser(&input); err != nil {
		s.logger.Errorw("Error parsing body", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, err.Error())
	}

	// получаем юзера из бд по имени
	user, err := s.repo.SingIn(ctx.Context(), input.Username)
	if err != nil {
		s.logger.Errorw("Error authorization", "error", err)
		return dto.InternalServerError(ctx)
	}

	// сверяем пароль из БД с введенныи паролем
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		s.logger.Errorw("Error wrong password", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, err.Error())
	}

	// создаем структуру для нашего jwt токена и передаем в него нужную информацию
	claims := CustomClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Subject:   strconv.Itoa(user.Id),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		UserId:   user.Id,
		Username: input.Username,
	}

	// создание токена
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// подписываем токен нашим ключем
	tokenString, err := token.SignedString([]byte(s.JWTSecret))
	if err != nil {
		s.logger.Errorw("Error signing token", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
		Data:   tokenString,
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// эта функция без middleware, просто функция для удаления пользователя по тз с ON DELETE CASCADE
func (s *UserService) DeleteUser(ctx *fiber.Ctx) error {
	id, err := strconv.Atoi(ctx.Params("id"))
	if err != nil {
		s.logger.Errorw("Error parsing id", "error", err)
		return dto.BadResponseError(ctx, dto.FieldBadFormat, err.Error())
	}

	err = s.repo.DeleteUser(ctx.Context(), id)
	if err != nil {
		s.logger.Errorw("Error deleting user", "error", err)
		return dto.InternalServerError(ctx)
	}

	response := dto.Response{
		Status: "success",
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
