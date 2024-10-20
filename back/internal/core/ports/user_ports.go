package ports

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	ListUsers(ctx context.Context, topic string) (*[]domain.User, *domain.ApiError)
	RegisterToNewsletter(ctx context.Context, email string, topics []string) *domain.ApiError
	UnregisterToNewsletter(ctx context.Context, email string, topic string) *domain.ApiError
}

type UserHandlers interface {
	ListUsers(c *fiber.Ctx) error
	RegisterToNewsletter(c *fiber.Ctx) error
	UnregisterToNewsletter(c *fiber.Ctx) error
}
