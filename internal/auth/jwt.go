package auth

import (
	"crypto/rand"     // Package untuk membuat token random aman secara kriptografi
	"crypto/sha256"   // Package untuk hash refresh token
	"encoding/base64" // Untuk encoding token agar bisa dikirim melalui JSON/HTTP
	"encoding/hex"    // Mengubah hash menjadi string
	"errors"
	"fmt"
	"time" // Untuk TTL (Time to Live) token

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	AccessSecret string        // Kunci rahasia untuk menandatangani access token (HS256)
	AccessTTL    time.Duration // Durasi aktif token
	RefreshTTL   time.Duration // Durasi aktif refresh token
}

// Fungsi dibawah adalah constuctor yang akan dipanggil di main

func NewJWTManager(secret string, accessTTL, refreshTTL time.Duration) *JWTManager {
	return &JWTManager{
		AccessSecret: secret,
		AccessTTL:    accessTTL,
		RefreshTTL:   refreshTTL,
	}
}

/*
Fungsi untuk generate Access token
*/

func (m *JWTManager) GenerateAccessToken(userID uint) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("%d", userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(m.AccessTTL)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(m.AccessSecret))
}

/*
Fungsi untuk verifikasi Access token
*/

func (m *JWTManager) VerifyAccessToken(tokenStr string) (*jwt.Token, error) {
	claims := &jwt.RegisteredClaims{}

	parser := jwt.NewParser(
		jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}),
		jwt.WithLeeway(0),
	)

	token, err := parser.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(m.AccessSecret), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("token tidak valid")
	}

	return token, nil
}

/*
Fungsi untuk generate Refresh token
*/

func (m *JWTManager) GenerateRefreshTokenRaw() (string, error) {
	// Membuat array 32 byte kosong
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	// Byte yang sudah dibuat di-encode menjadi string Base64 agar aman ditransfer via JSON
	return base64.URLEncoding.EncodeToString(b), nil
}

/*
Fungsi untuk hash Refresh token
*/

func (m *JWTManager) HashRefreshToken(raw string) string {
	// Menghasilkan array 32 byte hasil hash SHA256
	sum := sha256.Sum256([]byte(raw))
	// Mengubah hasil hash menjadi string hex
	return hex.EncodeToString(sum[:])
}
