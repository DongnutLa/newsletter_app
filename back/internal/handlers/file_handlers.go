package handlers

import (
	"bytes"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/gofiber/fiber/v2"
)

type FileHandlers struct {
	fileService ports.FilesService
}

// !TEST
var _ ports.FilesHandlers = (*FileHandlers)(nil)

func NewFileHandlers(fileService ports.FilesService) *FileHandlers {
	return &FileHandlers{
		fileService: fileService,
	}
}

func (f *FileHandlers) SaveFile(c *fiber.Ctx) error {
	folder := c.FormValue("folder")
	fileName := c.FormValue("fileName")
	if len(folder) == 0 {
		apiErr := domain.ErrInvalidParams.SetDetail("Folder is required")
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	buf := c.Locals(utils.FILE_KEY).(*bytes.Buffer)
	originalFileName := c.Locals(utils.FILE_NAME_KEY).(*string)
	if len(fileName) == 0 {
		fileName = *originalFileName
	}

	fileUrl, apiErr := f.fileService.SaveFile(c.Context(), buf, fileName, folder)
	if apiErr != nil {
		return c.Status(apiErr.HttpStatusCode).JSON(apiErr)
	}

	return c.Status(fiber.StatusOK).SendString(fileUrl)
}
