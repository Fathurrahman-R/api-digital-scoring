package auth

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type GinAuthMiddleware struct {
	jwt *JWTManager
}

func NewGinAuthMiddleware(jwt *JWTManager) *GinAuthMiddleware {
	return &GinAuthMiddleware{jwt}
}

func (m *GinAuthMiddleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		// Mengambil header authorization (Authorization: Bearer <access_token>)
		authHeader := c.GetHeader("Authorization")
		// Validasi format, harus memiliki header dan diawali dengan Bearer
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing or invalid Authorization header"})
			c.Abort()
			return
		}

		// Mengambil hanya raw tokennya saja
		rawToken := strings.TrimPrefix(authHeader, "Bearer ")

		// Verifikasikan token
		token, err := m.jwt.VerifyAccessToken(rawToken)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid or expired access token"})
			c.Abort()
			return
		}

		// Mengambil jwt claims dan disimpan sebagai map
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token claims"})
			c.Abort()
			return
		}

		// Ambil user id didalam claim dan diparsing dari JSON number ke float64
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user_id missing in token"})
			c.Abort()
			return
		}

		// Konversi float user id ke uint
		userID := uint(userIDFloat)

		// Simpan userID ke Gin Context untuk gunakan di handler
		c.Set("userID", userID)

		// Lanjutkan ke handler
		c.Next()
	}
}
