package ports

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	RegisterToNewsletter(ctx context.Context, email string) *domain.ApiError
	UnregisterToNewsletter(ctx context.Context, email string) *domain.ApiError
}

type UserHandlers interface {
	RegisterToNewsletter(c *fiber.Ctx) error
	UnregisterToNewsletter(c *fiber.Ctx) error
}
