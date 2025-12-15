.PHONY: help migrate migrate-fresh migrate-rollback migrate-status seed fresh-seed run build

# Default target
help:
	@echo "Available commands:"
	@echo "  make migrate         - Run database migrations"
	@echo "  make migrate-fresh   - Drop all tables and re-run migrations"
	@echo "  make migrate-rollback- Drop all tables"
	@echo "  make migrate-status  - Check migration status"
	@echo "  make seed            - Seed database with dummy data"
	@echo "  make fresh-seed      - Fresh migrate + seed"
	@echo "  make run             - Run the application"
	@echo "  make build           - Build the application"

# Run migrations
migrate:
	@echo "Running migrations..."
	@go run cmd/migrate/main.go migrate

# Fresh migration (drop all & recreate)
migrate-fresh:
	@echo "Running fresh migration..."
	@go run cmd/migrate/main.go migrate:fresh

# Rollback migrations
migrate-rollback:
	@echo "Rolling back migrations..."
	@go run cmd/migrate/main.go migrate:rollback

# Check migration status
migrate-status:
	@echo "Checking migration status..."
	@go run cmd/migrate/main.go migrate:status

# Seed database
seed:
	@echo "Seeding database..."
	@go run cmd/migrate/main.go db:seed

# Fresh migrate + seed
fresh-seed:
	@echo "Running fresh migration and seeding..."
	@go run cmd/migrate/main.go migrate:fresh-seed

# Run the application
run:
	@echo "Starting application..."
	@go run cmd/server/main.go

# Build the application
build:
	@echo "Building application..."
	@go build -o bin/server cmd/server/main.go
	@go build -o bin/migrate cmd/migrate/main.go
	@echo "Build complete! Binaries in ./bin/"

# Development: fresh migrate, seed, then run
dev: fresh-seed run