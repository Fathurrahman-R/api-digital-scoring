package auth

import (
	"api-digital-scoring/internal/entity"
	"context"
	"time"

	"gorm.io/gorm"
)

// Struct ini menyimpan pointer GORM DB
// Service layer hanya menerima interface RefreshTokenRepository
type refreshTokenRepo struct {
	db *gorm.DB
}

// NewRefreshTokenRepository : Constructor yang akan dipanggil di main

func NewRefreshTokenRepository(db *gorm.DB) RefreshTokenRepository {
	return &refreshTokenRepo{db}
}

// Create : Dipakai di login pertama dan refresh token rotation

func (r *refreshTokenRepo) Create(ctx context.Context, token *entity.RefreshToken) error {
	// Semua query memakai context
	return r.db.WithContext(ctx).Create(&token).Error
}

// GetByHash : Dipakai service untuk memastikan token tidak expired/revoked
func (r *refreshTokenRepo) GetByHash(ctx context.Context, hash string) (*entity.RefreshToken, error) {

	// Membuat variabel kosong untuk menampung hasil query
	var token entity.RefreshToken

	// Query berdasarkan hash dan ambil satu row
	err := r.db.WithContext(ctx).
		Where("token_hash = ?", hash).
		First(&token).Error

	if err != nil {
		return nil, err
	}
	return &token, nil
}

func (r *refreshTokenRepo) Revoke(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).
		// Tidak load data dulu tapi langsung update
		Model(&entity.RefreshToken{}).
		// Menunjukkan token mana yang akan direvoke
		Where("id = ?", id).
		// Set kolom revoked menjadi true
		Update("revoked", true).Error
}

// RevokeAllByUser : revoke semua token jika logout dari semua device
func (r *refreshTokenRepo) RevokeAllByUser(ctx context.Context, userID uint) error {
	return r.db.WithContext(ctx).
		Model(&entity.RefreshToken{}).
		Where("user_id = ?", userID).
		Update("revoked", true).Error
}

// DeleteExpired : Dilakukan secara berkala secara manual atau via cron job
func (r *refreshTokenRepo) DeleteExpired(ctx context.Context) error {
	return r.db.WithContext(ctx).
		Where("expires_at < ?", time.Now()).
		Delete(&entity.RefreshToken{}).Error
}
