package domain

import (
	"github.com/gofiber/fiber/v2"
)

var (
	ErrInvalidCredentials = NewApiError(
		"Invalid credentials",
		AuthErrors,
		fiber.StatusUnauthorized,
		1,
	)

	ErrUploadFile = NewApiError(
		"Failed to upload file",
		FileErrors,
		fiber.StatusInternalServerError,
		2,
	)
	ErrFetchFile = NewApiError(
		"Failed to fetch file",
		FileErrors,
		fiber.StatusInternalServerError,
		3,
	)

	ErrCreateNewsletter = NewApiError(
		"Failed to create newsletter",
		NewsletterErrors,
		fiber.StatusInternalServerError,
		4,
	)
	ErrNewsletterNotFound = NewApiError(
		"Newsletter not found",
		NewsletterErrors,
		fiber.StatusNotFound,
		5,
	)
	ErrSendEmailFailed = NewApiError(
		"Failed to send email",
		NewsletterErrors,
		fiber.StatusInternalServerError,
		6,
	)
	ErrFetchNewsletters = NewApiError(
		"Failed to fetch newsletters",
		NewsletterErrors,
		fiber.StatusInternalServerError,
		7,
	)

	ErrAddUserFailed = NewApiError(
		"Failed to add user to newsletters",
		UserErrors,
		fiber.StatusInternalServerError,
		8,
	)
	ErrDeleteUserFailed = NewApiError(
		"Failed to delete user to newsletters",
		UserErrors,
		fiber.StatusInternalServerError,
		9,
	)

	ErrFailedToParseBody = NewApiError(
		"Failed to parse body",
		GeneralErrors,
		fiber.StatusInternalServerError,
		10,
	)
	ErrInvalidParams = NewApiError(
		"Invalid params",
		GeneralErrors,
		fiber.StatusBadRequest,
		11,
	)
	ErrGenerateToken = NewApiError(
		"Unable to generate token access",
		GeneralErrors,
		fiber.StatusInternalServerError,
		12,
	)
)
