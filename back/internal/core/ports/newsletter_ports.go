package ports

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type NewsletterService interface {
	ListNewsletters(ctx context.Context, params *domain.PaginationsParams) (*domain.PaginatedResponse[domain.Newsletter], *domain.ApiError)
	CreateNewsletter(ctx context.Context, dto *domain.CreateNewsletterDTO) (*domain.Newsletter, *domain.ApiError)
	SendNewsletter(ctx context.Context, dto *domain.SendNewsletterDTO) *domain.ApiError
	ScheduleNewsletter(ctx context.Context) *domain.ApiError
}

type NewsletterHandlers interface {
	ListNewsletters(c *fiber.Ctx) error
	CreateNewsletter(c *fiber.Ctx) error
	SendNewsletter(c *fiber.Ctx) error
	ScheduleNewsletter(c *fiber.Ctx) error
}
