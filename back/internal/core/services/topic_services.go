package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/rs/zerolog"
)

type TopicService struct {
	logger    *zerolog.Logger
	topicRepo *repositories.TopicRepository
}

var _ ports.TopicService = (*TopicService)(nil)

func NewTopicService(ctx context.Context, logger *zerolog.Logger, repository *repositories.TopicRepository) *TopicService {
	return &TopicService{
		logger:    logger,
		topicRepo: repository,
	}
}

func (u *TopicService) ListTopics(ctx context.Context) (*[]domain.Topic, *domain.ApiError) {
	topics := []domain.Topic{}
	opts := ports.FindManyOpts{}

	_, err := u.topicRepo.Repo.FindMany(ctx, opts, &topics, false)
	if err != nil {
		u.logger.Error().Err(err).Msg("Failed to fetch topics")
		return nil, domain.ErrFetchTopic
	}

	return &topics, nil
}
