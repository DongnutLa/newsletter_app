package repositories

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type NewsletterRepository struct {
	Repo ports.Repository[domain.Newsletter, any]
}

func NewNewsletterRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) *NewsletterRepository {
	repo := BuildNewRepository[domain.Newsletter, any](ctx, collection, connection, logger)
	return &NewsletterRepository{
		Repo: repo,
	}
}
