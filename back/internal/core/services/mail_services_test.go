package services

import (
	"context"
	"os"
	"strconv"
	"testing"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	gomail "gopkg.in/mail.v2"
)

func TestSendEmails(t *testing.T) {
	assert := assert.New(t)

	utils.LoadConfig()
	host := utils.GetConfig("smtp_host")
	port := utils.GetConfig("smtp_port")
	user := utils.GetConfig("smtp_user")
	pwd := utils.GetConfig("smtp_pwd")
	intPort, _ := strconv.Atoi(port)

	// Dummies
	ctx := context.Background()
	logger := zerolog.New(os.Stderr)
	dialer := gomail.NewDialer(host, intPort, user, pwd)
	engine := html.New("./internal/template", ".html")

	service := NewMailService(ctx, &logger, dialer, engine)
	newsletter := domain.Newsletter{
		ID:         primitive.ObjectID{},
		Template:   "TEST",
		Subject:    "TEST",
		File:       "TEST",
		Recipients: []string{"dongnut09@gmail.com"},
		Topic:      "TEST",
	}
	payload := map[string]interface{}{
		"newsletter": &newsletter,
	}

	err := service.SendEmails(ctx, payload, string(domain.SendEmailTopic))

	assert.Nil(err, "Error must be nil")
}
