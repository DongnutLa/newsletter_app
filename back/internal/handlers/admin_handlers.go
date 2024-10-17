package handlers

import (
	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/gofiber/fiber/v2"
)

type AdminHandlers struct {
	adminService ports.AdminService
}

// !TEST
var _ ports.AdminHandlers = (*AdminHandlers)(nil)

func NewAdminHandlers(adminService ports.AdminService) *AdminHandlers {
	return &AdminHandlers{
		adminService: adminService,
	}
}

func (h *AdminHandlers) Login(c *fiber.Ctx) error {
	dto := domain.LoginDTO{}
	err := c.BodyParser(&dto)
	if err != nil {
		apiErr := domain.ErrFailedToParseBody
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	admin, apiErr := h.adminService.Login(c.Context(), &dto)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	c.Status(fiber.StatusOK).JSON(admin)
	return nil
}
