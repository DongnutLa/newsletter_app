package handlers

import (
	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type NewsletterHandlers struct {
	newsletterService ports.NewsletterService
}

// !TEST
var _ ports.NewsletterHandlers = (*NewsletterHandlers)(nil)

func NewNewsletterHandlers(newsletterService ports.NewsletterService) *NewsletterHandlers {
	return &NewsletterHandlers{
		newsletterService: newsletterService,
	}
}

func (n *NewsletterHandlers) CreateNewsletter(c *fiber.Ctx) error {
	dto := domain.CreateNewsletterDTO{}

	if err := c.BodyParser(&dto); err != nil {
		apiErr := domain.ErrFailedToParseBody
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	newsletter, apiErr := n.newsletterService.CreateNewsletter(c.Context(), &dto)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return c.Status(fiber.StatusCreated).JSON(newsletter)
}
func (n *NewsletterHandlers) SendNewsletter(c *fiber.Ctx) error {
	dto := domain.SendNewsletterDTO{}

	if err := c.BodyParser(&dto); err != nil {
		apiErr := domain.ErrFailedToParseBody
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	apiErr := n.newsletterService.SendNewsletter(c.Context(), &dto)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return c.Status(fiber.StatusOK).SendString("OK")
}

func (n *NewsletterHandlers) ScheduleNewsletter(c *fiber.Ctx) error {
	return nil
}
