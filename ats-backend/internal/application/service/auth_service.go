package service

import (
	"context"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"

	"ats-backend/internal/application/dto"
	"ats-backend/internal/application/port"
)

type authService struct {
	userRepo port.UserRepository
}

func NewAuthService(userRepo port.UserRepository) port.AuthService {
	return &authService{
		userRepo: userRepo,
	}
}

func (s *authService) Login(ctx context.Context, req *dto.LoginRequest) (*dto.LoginResponse, error) {
	// 1. Fetch user by email
	user, err := s.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 2. Check password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	// 3. Ensure user is active
	if !user.IsActive {
		return nil, errors.New("user account is inactive")
	}

	// 4. Generate JWT tokens
	accessToken, err := s.generateToken(user.ID.String(), user.TenantID.String(), user.Role.Name, 15*time.Minute)
	if err != nil {
		return nil, fmt.Errorf("failed to generate access token: %w", err)
	}

	refreshToken, err := s.generateToken(user.ID.String(), user.TenantID.String(), user.Role.Name, 7*24*time.Hour)
	if err != nil {
		return nil, fmt.Errorf("failed to generate refresh token: %w", err)
	}

	// 5. Construct response
	return &dto.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		User: dto.UserResponse{
			ID:        user.ID.String(),
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
			Role:      user.Role.Name,
			TenantID:  user.TenantID.String(),
		},
	}, nil
}

func (s *authService) generateToken(userID, tenantID, role string, expiration time.Duration) (string, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		jwtSecret = "super-secret-key-for-dev" // fallback for development
	}

	claims := jwt.MapClaims{
		"sub":    userID,
		"tenant": tenantID,
		"role":   role,
		"exp":    time.Now().Add(expiration).Unix(),
		"iat":    time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(jwtSecret))
}
