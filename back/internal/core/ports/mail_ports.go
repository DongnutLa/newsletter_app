package ports

import (
	"context"
)

type MailService interface {
	SendEmails(ctx context.Context, payload map[string]interface{}, topic string) error
}
