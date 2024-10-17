package repositories

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type AdminRepository struct {
	Repo ports.Repository[domain.Admin, any]
}

func NewAdminRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) *AdminRepository {
	repo := BuildNewRepository[domain.Admin, any](ctx, collection, connection, logger)
	return &AdminRepository{
		Repo: repo,
	}
}
