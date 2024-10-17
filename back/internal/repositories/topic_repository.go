package repositories

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type TopicRepository struct {
	Repo ports.Repository[domain.Topic, any]
}

func NewTopicRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) *TopicRepository {
	repo := BuildNewRepository[domain.Topic, any](ctx, collection, connection, logger)
	return &TopicRepository{
		Repo: repo,
	}
}
