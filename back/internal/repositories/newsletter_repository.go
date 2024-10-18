package repositories

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

//go:generate mockery --name=INewsletterRepository --inpackage=true
type INewsletterRepository interface {
	ports.Repository[domain.Newsletter, any]
}
type NewsletterRepository struct {
	ports.Repository[domain.Newsletter, any]
}

func NewNewsletterRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) INewsletterRepository {
	repo := BuildNewRepository[domain.Newsletter, any](ctx, collection, connection, logger)
	return &NewsletterRepository{
		repo,
	}
}
