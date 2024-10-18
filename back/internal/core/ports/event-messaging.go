package ports

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
)

//go:generate mockery --name=EventMessaging --inpackage=true
type EventMessaging interface {
	SendMessage(context.Context, *domain.MessageEvent)
}

type EventMessagingTypes string

const (
	// UseSNS EventMessagingTypes = "UseSNS"
	UseBUS EventMessagingTypes = "UseBUS"
)
