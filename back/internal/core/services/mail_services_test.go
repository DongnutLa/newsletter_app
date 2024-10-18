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
		File:       "https://upload.wikimedia.org/wikipedia/commons/0/07/%D0%A4%D0%9A_%22%D0%9A%D0%BE%D0%BB%D0%BE%D1%81%22_%28%D0%97%D0%B0%D1%87%D0%B5%D0%BF%D0%B8%D0%BB%D0%BE%D0%B2%D0%BA%D0%B0%2C_%D0%A5%D0%B0%D1%80%D1%8C%D0%BA%D0%BE%D0%B2%D1%81%D0%BA%D0%B0%D1%8F_%D0%BE%D0%B1%D0%BB%D0%B0%D1%81%D1%82%D1%8C%29_-_%D0%A4%D0%9A_%22%D0%91%D0%B0%D0%BB%D0%BA%D0%B0%D0%BD%D1%8B%22_%28%D0%97%D0%B0%D1%80%D1%8F%2C_%D0%9E%D0%B4%D0%B5%D1%81%D1%81%D0%BA%D0%B0%D1%8F_%D0%BE%D0%B1%D0%BB%D0%B0%D1%81%D1%82%D1%8C%29_%2818790931100%29.jpg",
		Recipients: []string{"dongnut09@gmail.com"},
		Topic:      "TEST",
	}
	payload := map[string]interface{}{
		"newsletter": &newsletter,
	}

	err := service.SendEmails(ctx, payload, string(domain.SendEmailTopic))

	assert.Nil(err, "Error must be nil")
}
