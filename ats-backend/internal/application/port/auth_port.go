package port

import (
	"context"

	"ats-backend/internal/application/dto"
)

type AuthService interface {
	Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error)
	// Additional methods like Register, RefreshToken, ValidateToken can be added here
}
