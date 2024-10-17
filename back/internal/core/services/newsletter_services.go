package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/rs/zerolog"
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

func (n *NewsletterService) ScheduleNewsletter(ctx context.Context) *domain.ApiError {
	//
	return nil
}
