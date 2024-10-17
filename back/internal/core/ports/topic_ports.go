package ports

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type TopicService interface {
	ListTopics(ctx context.Context) (*[]domain.Topic, *domain.ApiError)
}

type TopicHandlers interface {
	ListTopics(c *fiber.Ctx) error
}
