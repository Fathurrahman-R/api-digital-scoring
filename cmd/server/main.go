package main

import (
	"api-digital-scoring/internal/auth"
	"api-digital-scoring/internal/config"
	"api-digital-scoring/internal/user"
	"api-digital-scoring/migrations"
	"api-digital-scoring/pkg/database"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

var accessTTL time.Duration
var refreshTTL time.Duration

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to Load config: %v", err)
	}

	// initializing database connection
	db, err := database.NewMySQL(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection established")

	if err = migrations.Migrate(db); err != nil {
		panic(err)
	}

	// JWT contruct //
	accessTTL, err := time.ParseDuration(cfg.JWT.AccessTokenLifetime)
	if err != nil {
		log.Fatalf("invalid access_token_lifetime: %v", err)
	}

	refreshTTL, err := time.ParseDuration(cfg.JWT.RefreshTokenLifetime)
	if err != nil {
		log.Fatalf("invalid refresh_token_lifetime: %v", err)
	}
	jwtManager := auth.NewJWTManager(cfg.JWT.AccessSecret, accessTTL, refreshTTL)

	// Repositories //
	rtRepo := auth.NewRefreshTokenRepository(db)
	uRepo := user.NewUserRepo(db)

	// Services //
	authService := auth.NewAuthService(rtRepo, jwtManager, uRepo)
	userService := user.NewService(uRepo)

	// Handlers //
	authHandler := auth.NewHandler(authService)
	userHandler := user.NewHandler(userService)

	// Middlewares //
	authMiddleware := auth.NewGinAuthMiddleware(jwtManager)

	r := gin.Default()

	// Routes //
	authGroup := r.Group("/auth")
	authHandler.RegisterRoutes(authGroup)

	private := r.Group("/")
	private.Use(authMiddleware.Middleware())
	userHandler.RegisterRoutes(private)

	err = r.Run(":" + cfg.Server.Port)
	if err != nil {
		return
	}
}
