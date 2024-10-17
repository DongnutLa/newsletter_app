package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/rs/zerolog"
	"github.com/samber/lo"
)

type UserService struct {
	logger    *zerolog.Logger
	userRepo  *repositories.UserRepository
	messaging ports.EventMessaging
}

var _ ports.UserService = (*UserService)(nil)

func NewUserService(
	ctx context.Context,
	logger *zerolog.Logger,
	repository *repositories.UserRepository,
	messaging ports.EventMessaging,
) *UserService {
	return &UserService{
		logger:    logger,
		userRepo:  repository,
		messaging: messaging,
	}
}

func (u *UserService) ListUsers(ctx context.Context, topic string) (*[]domain.User, *domain.ApiError) {
	users := []domain.User{}

	filter := map[string]interface{}{}
	if topic != "" {
		filter["topics"] = topic
	}

	opts := ports.FindManyOpts{
		Filter: filter,
	}

	_, err := u.userRepo.Repo.FindMany(ctx, opts, &users, false)
	if err != nil {
		u.logger.Error().Err(err).Msg("Failed to fetch users")
		return nil, domain.ErrFetchUser
	}

	return &users, nil
}

func (u *UserService) RegisterToNewsletter(ctx context.Context, email string, topics []string) *domain.ApiError {
	foundUser := domain.User{}
	findOpts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"email": email,
		},
	}
	if err := u.userRepo.Repo.FindOne(ctx, findOpts, &foundUser); err != nil && err.Error() != "not found" {
		u.logger.Error().Err(err).Msg("Failed finding user to add to newsletters")
		return domain.ErrAddUserFailed
	}

	if foundUser.Timestamp != nil {
		updOpts := ports.UpdateOpts{
			Filter: map[string]interface{}{
				"email": email,
			},
			Payload: &map[string]interface{}{
				"topics": topics,
			},
		}

		_, err := u.userRepo.Repo.UpdateOne(ctx, updOpts)
		if err != nil {
			u.logger.Error().Err(err).Msg("Failed updating user to add to newsletters")
			return domain.ErrAddUserFailed
		}
	} else {
		newUser := domain.NewUser(email, topics)
		if err := u.userRepo.Repo.InsertOne(ctx, *newUser); err != nil {
			u.logger.Error().Err(err).Msg("Failed to add user to newsletters")
			return domain.ErrAddUserFailed
		}
	}

	return nil
}

func (u *UserService) UnregisterToNewsletter(ctx context.Context, email string, topic string) *domain.ApiError {
	user := domain.User{}
	opts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"email": email,
		},
	}
	if err := u.userRepo.Repo.FindOne(ctx, opts, &user); err != nil {
		u.logger.Error().Err(err).Msg("Failed to find user to unregister")
		return domain.ErrDeleteUserFailed
	}

	newTopics := lo.Filter(user.Topics, func(item string, _ int) bool {
		return item != topic
	})

	updtOpts := ports.UpdateOpts{
		Filter: map[string]interface{}{
			"email": email,
		},
		Payload: &map[string]interface{}{
			"topics": newTopics,
		},
	}

	_, err := u.userRepo.Repo.UpdateOne(ctx, updtOpts)
	if err != nil {
		u.logger.Error().Err(err).Msg("Failed to delete topic for user from newsletters")
		return domain.ErrDeleteUserFailed
	}

	//Propagate changes in newsletters
	evt := domain.MessageEvent{
		EventTopic: domain.PropagateUserUnsubscription,
		Data: map[string]interface{}{
			"email": email,
			"topic": topic,
		},
	}
	u.messaging.SendMessage(ctx, &evt)

	return nil
}
