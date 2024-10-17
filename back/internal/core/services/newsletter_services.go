package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type NewsletterService struct {
	logger         *zerolog.Logger
	newsletterRepo *repositories.NewsletterRepository
	messaging      ports.EventMessaging
}

var _ ports.NewsletterService = (*NewsletterService)(nil)

func NewNewsletterService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository *repositories.NewsletterRepository,
	messaging ports.EventMessaging,
) *NewsletterService {
	return &NewsletterService{
		logger:         logger,
		newsletterRepo: repository,
		messaging:      messaging,
	}
}

func (n *NewsletterService) ListNewsletters(ctx context.Context, params *domain.PaginationsParams) (*domain.PaginatedResponse[domain.Newsletter], *domain.ApiError) {
	result := []domain.Newsletter{}

	opts := ports.FindManyOpts{
		Take: params.PageSize,
		Skip: params.PageSize * (params.Page - 1),
	}
	total, err := n.newsletterRepo.Repo.FindMany(ctx, opts, &result, true)
	if err != nil {
		return nil, domain.ErrFetchNewsletters
	}

	response := domain.PaginatedResponse[domain.Newsletter]{
		Metadata: domain.Pagination{
			Page:     params.Page,
			PageSize: params.PageSize,
			HasNext:  float64(params.Page) < (float64(*total) / float64(params.PageSize)),
			Length:   *total,
		},
		Data: result,
	}

	return &response, nil
}

func (n *NewsletterService) CreateNewsletter(ctx context.Context, dto *domain.CreateNewsletterDTO) (*domain.Newsletter, *domain.ApiError) {
	newsletter := domain.NewNewsletter(dto)

	if err := n.newsletterRepo.Repo.InsertOne(ctx, *newsletter); err != nil {
		return nil, domain.ErrCreateNewsletter
	}

	return newsletter, nil
}

func (n *NewsletterService) SendNewsletter(ctx context.Context, dto *domain.SendNewsletterDTO) *domain.ApiError {
	newsletter := domain.Newsletter{}

	opts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"_id": dto.NewsletterId,
		},
	}
	if err := n.newsletterRepo.Repo.FindOne(ctx, opts, &newsletter); err != nil {
		return domain.ErrNewsletterNotFound
	}

	// Send email event
	evt := domain.MessageEvent{
		EventTopic: domain.SendEmailTopic,
		Data: map[string]interface{}{
			"newsletter": &newsletter,
		},
	}
	n.messaging.SendMessage(ctx, &evt)

	return nil
}

func (n *NewsletterService) UnregisterUserFromNewsletter(ctx context.Context, payload map[string]interface{}, topic string) error {
	n.logger.Info().Interface("payload", payload).Msgf("Message received for topic %s", topic)

	newsTopic := payload["topic"].(string)
	email := payload["email"].(string)

	newsletter := domain.Newsletter{}
	opts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"topic": newsTopic,
		},
	}
	if err := n.newsletterRepo.Repo.FindOne(ctx, opts, &newsletter); err != nil {
		n.logger.Error().Err(err).Msgf("Newsletter for topic %s not found", newsTopic)
		return err
	}

	newRecipients := lo.Filter(newsletter.Recipients, func(recipient string, _ int) bool {
		return email != recipient
	})

	updOpts := ports.UpdateOpts{
		Filter: map[string]interface{}{
			"topic": newsTopic,
		},
		Payload: &map[string]interface{}{
			"recipients": newRecipients,
		},
	}

	_, err := n.newsletterRepo.Repo.UpdateOne(ctx, updOpts)
	if err != nil {
		n.logger.Error().Err(err).Msgf("Failed to update newsletter for topic %s", newsTopic)
		return err
	}

	n.logger.Error().Err(err).Msgf("Successfully updated newsletter for topic %s and user %s", newsTopic, email)

	return nil
}

func (n *NewsletterService) ScheduleNewsletter(ctx context.Context) *domain.ApiError {
	//
	return nil
}
