package auth

import (
	"api-digital-scoring/internal/entity"
	"context"
)

type RefreshTokenRepository interface {
	// Create new refresh token to DB setelah user login atau refresh
	Create(ctx context.Context, token *entity.RefreshToken) error

	// GetByHash mencari refresh token berdasarkan hash
	// Dipakai untuk: Validasi refresh token dan memastikan token masih aktif
	GetByHash(ctx context.Context, hash string) (*entity.RefreshToken, error)

	// Revoke satu token
	Revoke(ctx context.Context, id uint) error

	// RevokeAllByUser : revoke/logout dari semua device
	RevokeAllByUser(ctx context.Context, userID uint) error

	// DeleteExpired : membersihkan token yang sudah kedaluarsa
	DeleteExpired(ctx context.Context) error
}
