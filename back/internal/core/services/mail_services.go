package services

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/gofiber/template/html/v2"
	"github.com/rs/zerolog"
	"github.com/tdewolff/minify"
	mHtml "github.com/tdewolff/minify/html"
	gomail "gopkg.in/mail.v2"
)

type MailService struct {
	logger     *zerolog.Logger
	mailDialer *gomail.Dialer
	mail       string
	baseUrl    string
	engine     *html.Engine
}

var _ ports.MailService = (*MailService)(nil)

func NewMailService(
	ctx context.Context,
	logger *zerolog.Logger,
	dialer *gomail.Dialer,
	engine *html.Engine,
) *MailService {
	mail := utils.GetConfig("smtp_mail")
	baseUrl := utils.GetConfig("base_url")

	return &MailService{
		logger:     logger,
		mailDialer: dialer,
		mail:       mail,
		baseUrl:    baseUrl,
		engine:     engine,
	}
}

func (m *MailService) SendEmails(ctx context.Context, payload map[string]interface{}, topic string) error {
	m.logger.Info().Interface("payload", payload).Msgf("Message received for topic %s", topic)

	newsletter := utils.EventDataToStruct[domain.Newsletter](payload["newsletter"])

	var wg sync.WaitGroup

	for _, recipient := range newsletter.Recipients {
		wg.Add(1)
		go m.sendEmail(recipient, newsletter, &wg)
	}
	wg.Wait()
	return nil
}

func (m *MailService) sendEmail(recipient string, newsletter *domain.Newsletter, wg *sync.WaitGroup) error {
	defer wg.Done()

	url := fmt.Sprintf("%s/v1/users/unregister?email=%s&topic=%s", m.baseUrl, recipient, newsletter.Topic)

	buf := new(bytes.Buffer)
	err := m.engine.Render(buf, "template", map[string]string{
		"UnsubscribeUrl": url,
		"Template":       newsletter.Template,
	})
	if err != nil {
		m.logger.Err(err).Msg("Failed rendering template")
		return err
	}

	mediaType := "text/html"
	mn := minify.New()
	mn.AddFunc(mediaType, mHtml.Minify)
	html, err := mn.String(mediaType, buf.String())
	if err != nil {
		m.logger.Err(err).Msg("Failed rendering minified template")
		return err
	}

	message := gomail.NewMessage()
	// Set email headers
	message.SetHeader("From", m.mail)
	message.SetHeader("To", recipient)
	message.SetHeader("Subject", newsletter.Subject)
	message.SetBody("text/html", html)

	file, ext, err := downloadFile(newsletter.File)
	if err != nil {
		m.logger.Error().Err(err).Msg("Failed to use image for email")
		return domain.ErrFetchFile
	}
	defer file.Close()
	message.AttachReader(fmt.Sprintf("newsletter.%s", ext), file)
	if ext == "png" {
		message.EmbedReader("newsletter.png", file)
	}

	if err := m.mailDialer.DialAndSend(message); err != nil {
		m.logger.Error().Err(err).Msg("Failed to send email")
		return domain.ErrSendEmailFailed
	} else {
		m.logger.Info().Msg("Email sent successfully")
	}

	return nil
}

func downloadFile(url string) (io.ReadCloser, string, error) {
	// Realiza la solicitud GET a la URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, "", fmt.Errorf("error realizando solicitud GET: %v", err)
	}
	// defer resp.Body.Close()

	// Verifica que la solicitud sea exitosa (código de estado 200)
	if resp.StatusCode != http.StatusOK {
		return nil, "", fmt.Errorf("error en la respuesta del servidor: %v", resp.Status)
	}

	splitted := strings.Split(url, ".")
	ext := splitted[len(splitted)-1]

	// Puedes retornar el io.Reader (resp.Body) o leer los datos desde ahí
	return resp.Body, ext, nil
}
