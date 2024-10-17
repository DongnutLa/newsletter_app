package repositories

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/rs/zerolog"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	Repo ports.Repository[domain.User, any]
}

func NewUserRepository(ctx context.Context, collection string, connection *mongo.Database, logger *zerolog.Logger) *UserRepository {
	repo := BuildNewRepository[domain.User, any](ctx, collection, connection, logger)
	return &UserRepository{
		Repo: repo,
	}
}
