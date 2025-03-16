package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"pract2/internal/dto"
	"strings"
)

type CustomClaims struct {
	jwt.RegisteredClaims
	UserID   int    `json:"user_id"`
	Username string `json:"username"`
}

// будущее middleware

func Authorization(jwtToken string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return dto.UnauthorizedError(ctx)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return dto.UnauthorizedError(ctx)
		}

		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtToken), nil
		})

		if err != nil || !token.Valid {
			return dto.UnauthorizedError(ctx)
		}

		ctx.Locals("user_id", claims.UserID)
		ctx.Locals("username", claims.Username)

		return ctx.Next()
	}
}
