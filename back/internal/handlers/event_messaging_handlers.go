package handlers

import (
	"context"
	"fmt"
	"sync"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/services"
	"github.com/mustafaturan/bus/v3"
	"github.com/rs/zerolog"
)

var mailEvent chan bus.Event
var cancel context.CancelFunc
var c context.Context

var mailWorker = "mails"

type EventsHandler struct {
	MailService *services.MailService
	Logger      *zerolog.Logger
}

func NewEventsHandler(ctx context.Context, log *zerolog.Logger, mailService *services.MailService) *EventsHandler {
	return &EventsHandler{
		MailService: mailService,
		Logger:      log,
	}
}

func (h *EventsHandler) Start(wg *sync.WaitGroup) {
	c, cancel = context.WithCancel(context.Background())

	mailEvent = make(chan bus.Event)

	// Handlers
	mailHandler := bus.Handler{Handle: func(_ context.Context, e bus.Event) {
		mailEvent <- e
	}, Matcher: string(domain.SendEmailTopic)}

	services.Bus.RegisterHandler(mailWorker, mailHandler)

	fmt.Printf("Registered handlers...\n")

	wg.Add(4)
	go h.handler(wg)
}

// Stop deregisters handlers
func (h *EventsHandler) Stop() {
	defer fmt.Printf("Deregistered handlers...\n")
	services.Bus.DeregisterHandler(mailWorker)
	cancel()
}

func (h *EventsHandler) handler(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case <-c.Done():
			return
		case e := <-mailEvent:
			h.MailService.SendEmails(c, e.Data.(map[string]interface{}), e.Topic)
		}
	}
}
