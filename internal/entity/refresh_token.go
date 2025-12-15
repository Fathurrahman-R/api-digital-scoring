package entity

import "time"

type RefreshToken struct {
	ID        uint      `gorm:"primary_key"`
	UserID    uint      `gorm:"index;not null"`
	TokenHash string    `gorm:"uniqueIndex;not null; size:255"`
	ExpiresAt time.Time `gorm:"index;not null"`
	Revoked   bool      `gorm:"default:false"` // Untuk cek apakah token sudah di blokir/revoke
	CreatedAt time.Time `gorm:"index;not null"`
}
