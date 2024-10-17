package services

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/DongnutLa/newsletter_app/internal/utils"
	"github.com/rs/zerolog"
	gomail "gopkg.in/mail.v2"
)

type NewsletterService struct {
	logger         *zerolog.Logger
	newsletterRepo *repositories.NewsletterRepository
	mailDialer     *gomail.Dialer
	mail           string
}

var _ ports.NewsletterService = (*NewsletterService)(nil)

func NewNewsletterService(ctx context.Context, logger *zerolog.Logger, repository *repositories.NewsletterRepository, dialer *gomail.Dialer) *NewsletterService {
	mail := utils.GetConfig("smtp_mail")
	return &NewsletterService{
		logger:         logger,
		newsletterRepo: repository,
		mailDialer:     dialer,
		mail:           mail,
	}
}

func (n *NewsletterService) CreateNewsletter(ctx context.Context, dto *domain.CreateNewsletterDTO) (*domain.Newsletter, *domain.ApiError) {
	newsletter := domain.NewNewsletter(dto.Template, dto.File, dto.Recipients)

	if err := n.newsletterRepo.Repo.InsertOne(ctx, *newsletter); err != nil {
		return nil, domain.ErrCreateNewsletter
	}

	return nil, nil
}

func (n *NewsletterService) SendNewsletter(ctx context.Context, dto *domain.SendNewsletterDTO) *domain.ApiError {
	message := gomail.NewMessage()

	newsletter := domain.Newsletter{}

	opts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"_id": dto.NewsletterId,
		},
	}
	if err := n.newsletterRepo.Repo.FindOne(ctx, opts, &newsletter); err != nil {
		return domain.ErrNewsletterNotFound
	}

	recipients := append(newsletter.Recipients, dto.ExtraEmail)

	// Set email headers
	message.SetHeader("From", n.mail)
	message.SetHeader("To", recipients...)
	message.SetHeader("Subject", newsletter.Subject)

	//Attatch file
	template := fmt.Sprintf(`
		%s
		<p>Unsubscribe here</p>
	`, newsletter.Template)
	message.SetBody("text/html", template)

	file, err := downloadFile(newsletter.File)
	if err != nil {
		n.logger.Error().Err(err).Msg("Failed to use image for email")
		return domain.ErrFetchFile
	}
	message.AttachReader("newsletter", file)

	if err := n.mailDialer.DialAndSend(message); err != nil {
		n.logger.Error().Err(err).Msg("Failed to send email")
		return domain.ErrSendEmailFailed
	} else {
		n.logger.Info().Msg("Email sent successfully")
	}

	return nil
}

func (n *NewsletterService) ScheduleNewsletter(ctx context.Context) *domain.ApiError {
	//
	return nil
}

func downloadFile(url string) (io.Reader, error) {
	// Realiza la solicitud GET a la URL
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("error realizando solicitud GET: %v", err)
	}
	defer resp.Body.Close()

	// Verifica que la solicitud sea exitosa (código de estado 200)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error en la respuesta del servidor: %v", resp.Status)
	}

	// Puedes retornar el io.Reader (resp.Body) o leer los datos desde ahí
	return resp.Body, nil
}
