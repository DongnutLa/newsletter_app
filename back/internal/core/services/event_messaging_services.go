package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/mustafaturan/bus/v3"
	"github.com/rs/zerolog"
)

type BusMessaging struct {
	busMsg *bus.Bus
	logger *zerolog.Logger
}

func NewEventMessaging(
	log *zerolog.Logger,
	eventType ports.EventMessagingTypes,
) ports.EventMessaging {
	if eventType == ports.UseBUS {
		return NewBusMessaging(log)
	}

	/* if eventType == UseSNS {
		return NewSnsMessaging(log)
	} */

	return nil
}

func NewBusMessaging(logger *zerolog.Logger) ports.EventMessaging {
	log := logger.With().Str("resource", "bus_messaging").Logger()

	return &BusMessaging{
		busMsg: Bus,
		logger: &log,
	}
}

func (b *BusMessaging) SendMessage(ctx context.Context, data *domain.MessageEvent) {
	err := b.busMsg.Emit(ctx, string(data.EventTopic), data.Data)
	if err != nil {
		b.logger.Err(err).Interface("data", data).Msg("<BUS MESSAGING> Faild to send message")
	} else {
		b.logger.Log().Interface("data", data).Msg("<BUS MESSAGING> Message published successfully")
	}
}
