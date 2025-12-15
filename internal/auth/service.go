package auth

import (
	"api-digital-scoring/internal/entity"
	"api-digital-scoring/internal/user"
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrInvalidCredentials = errors.New("invalid username or password")
	ErrInvalidToken       = errors.New("invalid refresh token")
	ErrTokenExpired       = errors.New("refresh token expired")
	ErrTokenRevoked       = errors.New("refresh token revoked")
)

type Service struct {
	repo     RefreshTokenRepository // Akses ke tabel refresh token
	jwt      *JWTManager            // Memanggil fungsi-fungsi
	userRepo user.Repository        // Akses ke tabel user
	now      func() time.Time       // Mendapatkan waktu
}

// NewAuthService : Constuctor
func NewAuthService(repo RefreshTokenRepository, jwt *JWTManager, userRepo user.Repository) *Service {
	return &Service{
		repo:     repo,     // Dependency Injection
		jwt:      jwt,      // Menghasilkan token
		userRepo: userRepo, // Pencarian user
		now:      time.Now,
	}
}

func (s *Service) Login(ctx context.Context, username, password string) (string, string, error) {
	u, err := s.userRepo.GetByUsername(ctx, username)
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	if err != nil {
		return "", "", ErrInvalidCredentials
	}

	accessToken, err := s.jwt.GenerateAccessToken(u.ID)
	if err != nil {
		return "", "", err
	}

	rawRefresh, err := s.jwt.GenerateRefreshTokenRaw()
	if err != nil {
		return "", "", err
	}

	hash := s.jwt.HashRefreshToken(rawRefresh)

	rt := &entity.RefreshToken{
		UserID:    u.ID,
		TokenHash: hash,
		ExpiresAt: s.now().Add(s.jwt.RefreshTTL),
		Revoked:   false,
	}

	if err := s.repo.Create(ctx, rt); err != nil {
		return "", "", err
	}

	return accessToken, rawRefresh, nil // Kembalikan ke user, refresh token simpan di cookie
}

func (s *Service) Refresh(ctx context.Context, rawRefresh string) (string, string, error) {
	hash := s.jwt.HashRefreshToken(rawRefresh)

	token, err := s.repo.GetByHash(ctx, hash)
	if err != nil {
		return "", "", ErrInvalidToken
	}

	if token.Revoked {
		return "", "", ErrTokenRevoked
	}

	if token.ExpiresAt.Before(s.now()) {
		return "", "", ErrTokenExpired
	}

	newAccess, err := s.jwt.GenerateAccessToken(token.UserID)
	if err != nil {
		return "", "", err
	}

	newRawRefresh, err := s.jwt.GenerateRefreshTokenRaw()
	if err != nil {
		return "", "", err
	}

	newHash := s.jwt.HashRefreshToken(newRawRefresh)

	newToken := &entity.RefreshToken{
		UserID:    token.UserID,
		TokenHash: newHash,
		ExpiresAt: s.now().Add(s.jwt.RefreshTTL),
		Revoked:   false,
	}

	if err := s.repo.Create(ctx, newToken); err != nil {
		return "", "", err
	}

	_ = s.repo.Revoke(ctx, token.ID)

	return newAccess, newRawRefresh, nil
}

func (s *Service) Logout(ctx context.Context, rawRefresh string) error {
	hash := s.jwt.HashRefreshToken(rawRefresh)

	token, err := s.repo.GetByHash(ctx, hash)
	if err != nil {
		return ErrInvalidToken
	}

	return s.repo.Revoke(ctx, token.ID)
}
