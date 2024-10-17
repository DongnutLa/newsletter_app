package services

import (
	"context"

	"github.com/DongnutLa/newsletter_app/internal/core/domain"
	"github.com/DongnutLa/newsletter_app/internal/core/ports"
	"github.com/DongnutLa/newsletter_app/internal/repositories"
	"github.com/rs/zerolog"
)

type AdminService struct {
	logger     *zerolog.Logger
	adminRepo  *repositories.AdminRepository
	jwtService *JwtService
}

var _ ports.AdminService = (*AdminService)(nil)

func NewAdminService(ctx context.Context, logger *zerolog.Logger, repository *repositories.AdminRepository, jwtService *JwtService) *AdminService {
	return &AdminService{
		logger:     logger,
		adminRepo:  repository,
		jwtService: jwtService,
	}
}

func (a *AdminService) Login(ctx context.Context, dto *domain.LoginDTO) (*domain.Admin, *domain.ApiError) {
	opts := ports.FindOneOpts{
		Filter: map[string]interface{}{
			"email": dto.Email,
		},
	}

	admin := domain.Admin{}

	err := a.adminRepo.Repo.FindOne(ctx, opts, &admin)
	if err != nil {
		a.logger.Error().Err(err).Msgf("Failed to find user with email %s", dto.Email)
		return nil, domain.ErrInvalidCredentials
	}

	if admin.Password != dto.Password {
		return nil, domain.ErrInvalidCredentials
	}

	token, err := a.jwtService.GenerateJWT(admin.ID.Hex(), admin.Email, admin.Name)
	admin.Token = token

	return &admin, nil
}
