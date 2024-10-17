package ports

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type AdminService interface {
	Login(ctx context.Context, dto *domain.LoginDTO) (*domain.Admin, *domain.ApiError)
}

type AdminHandlers interface {
	Login(c *fiber.Ctx) error
}
