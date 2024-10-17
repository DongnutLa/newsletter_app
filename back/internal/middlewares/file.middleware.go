package middlewares

import (
	"bytes"
	"io"
	"slices"
	"strings"

	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

func NewFileMiddleware(logger *zerolog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		valid := []string{"image/png", "application/pdf"}

		head, err := c.FormFile("file")
		if err != nil {
			logger.Err(err).Msg("ERROR | Upload controller form file")
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		validType := slices.Contains(valid, head.Header.Get("Content-Type"))
		if !validType {
			return c.Status(fiber.StatusBadRequest).JSON("Invalid file type")
		}

		if head.Size > 4000000 {
			return c.Status(fiber.StatusBadRequest).JSON("File size over 2MB")
		}

		name := strings.Split(head.Filename, ".")[0]

		file, err := head.Open()
		if err != nil {
			logger.Err(err).Msg("ERROR | Upload controller open file")
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}
		defer file.Close()

		buf := bytes.NewBuffer(nil)
		if _, err = io.Copy(buf, file); err != nil {
			logger.Err(err).Msg("ERROR | Upload controller buffering file")
			return c.Status(fiber.StatusInternalServerError).JSON(err)
		}

		c.Locals(utils.FILE_KEY, buf)
		c.Locals(utils.FILE_NAME_KEY, &name)

		return c.Next()
	}
}
