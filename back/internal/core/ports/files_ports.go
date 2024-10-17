package ports

import (
	"bytes"
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/gofiber/fiber/v2"
)

type FilesService interface {
	SaveFile(ctx context.Context, buf *bytes.Buffer, fileName, folder string) (string, *domain.ApiError)
}

type FilesHandlers interface {
	SaveFile(c *fiber.Ctx) error
}
