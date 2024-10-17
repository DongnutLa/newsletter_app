package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/rs/zerolog"
)

type UserService struct {
	logger   *zerolog.Logger
	userRepo *repositories.UserRepository
}

var _ ports.UserService = (*UserService)(nil)

func NewUserService(ctx context.Context, logger *zerolog.Logger, repository *repositories.UserRepository) *UserService {
	return &UserService{
		logger:   logger,
		userRepo: repository,
	}
}

func (u *UserService) ListUsers(ctx context.Context) (*[]domain.User, *domain.ApiError) {
	users := []domain.User{}
	opts := ports.FindManyOpts{}

	_, err := u.userRepo.Repo.FindMany(ctx, opts, &users, false)
	if err != nil {
		u.logger.Error().Err(err).Msg("Failed to fetch users")
		return nil, domain.ErrFetchUser
	}

	return &users, nil
}

func (u *UserService) RegisterToNewsletter(ctx context.Context, email string) *domain.ApiError {
	newUser := domain.NewUser(email)

	if err := u.userRepo.Repo.InsertOne(ctx, *newUser); err != nil {
		u.logger.Error().Err(err).Msg("Failed to add user to newsletters")
		return domain.ErrAddUserFailed
	}

	return nil
}

func (u *UserService) UnregisterToNewsletter(ctx context.Context, email string) *domain.ApiError {
	opts := ports.DeleteOpts{
		Filter: map[string]interface{}{
			"email": email,
		},
	}

	ok, err := u.userRepo.Repo.DeleteOne(ctx, opts)
	if err != nil || !ok {
		u.logger.Error().Err(err).Msg("Failed to delete user from newsletters")
		return domain.ErrDeleteUserFailed
	}

	return nil
}
