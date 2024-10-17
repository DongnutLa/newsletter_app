package middlewares

import (
	"strings"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/services"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func NewAuthMiddleware(logger *zerolog.Logger, jwtService *services.JwtService) fiber.Handler {
	return func(c *fiber.Ctx) error {
		authHeader := c.Get("Authorization")
		if authHeader == "" {
			apiErr := domain.ErrInvalidCredentials
			return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		claims, apiErr := jwtService.VerifyJWT(tokenString)
		if apiErr != nil {
			return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}
		if claims == nil || claims.ID == "" {
			apiErr := domain.ErrInvalidCredentials
			return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
		}

		return c.Next()
	}
}
