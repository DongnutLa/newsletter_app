package handlers

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type TopicHandlers struct {
	topicService ports.TopicService
}

// !TEST
var _ ports.TopicHandlers = (*TopicHandlers)(nil)

func NewTopicHandlers(topicService ports.TopicService) *TopicHandlers {
	return &TopicHandlers{
		topicService: topicService,
	}
}

func (h *TopicHandlers) ListTopics(c *fiber.Ctx) error {
	topics, apiErr := h.topicService.ListTopics(context.TODO())
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return c.Status(fiber.StatusOK).JSON(topics)
}
