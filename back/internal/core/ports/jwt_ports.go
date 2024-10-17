package ports

import "github.com/DongnutLa/newsletter_app/internal/core/domain"

type JwtService interface {
	GenerateJWT(id, email, name string) (string, *domain.ApiError)
	VerifyJWT(tokenString string) (*domain.Claims, *domain.ApiError)
}
