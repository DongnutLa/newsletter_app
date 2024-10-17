package handlers

import (
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	userService ports.UserService
}

// !TEST
var _ ports.UserHandlers = (*UserHandlers)(nil)

func NewUserHandlers(userService ports.UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

func (h *UserHandlers) RegisterToNewsletter(c *fiber.Ctx) error {
	email := c.Query("email", "")

	apiErr := h.userService.RegisterToNewsletter(c.Context(), email)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	c.Status(fiber.StatusOK).SendString("OK")

	return nil
}

func (h *UserHandlers) UnregisterToNewsletter(c *fiber.Ctx) error {
	email := c.Query("email", "")

	apiErr := h.userService.UnregisterToNewsletter(c.Context(), email)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	c.Status(fiber.StatusOK).SendString("OK")
	return nil
}
