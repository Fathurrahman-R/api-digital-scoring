package main

import (
	"api-digital-scoring/internal/config"
	"api-digital-scoring/migrations"
	"api-digital-scoring/pkg/database"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func main() {

	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to Load config: %v", err)
	}

	// Initialize database connection
	db, err := database.NewMySQL(cfg.Database)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Get SQL DB untuk close connection
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// Close connection after finishing function execution
	defer func(sqlDB *sql.DB) {
		_ = sqlDB.Close()
	}(sqlDB)

	// Cek command line argument
	if len(os.Args) < 2 {
		printUsage()
		os.Exit(1)
	}

	command := os.Args[1]

	switch command {
	case "migrate":
		// Migrasi biasa
		if err = migrations.Migrate(db); err != nil {
			fmt.Printf("❌ Migration failed: %v\n", err)
			os.Exit(1)
		}

	case "migrate:fresh":
		// Fresh migration (drop all & recreate)
		fmt.Println("⚠️  WARNING: This will DROP ALL TABLES and data!")
		fmt.Print("Are you sure? (yes/no): ")

		var confirm string
		_, err = fmt.Scanln(&confirm)
		if err != nil {
			return
		}

		if confirm == "yes" {
			if err = migrations.FreshMigrate(db); err != nil {
				fmt.Printf("❌ Fresh migration failed: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("❌ Migration cancelled")
		}

	case "migrate:rollback":
		// Rollback (drop all tables)
		fmt.Println("⚠️  WARNING: This will DROP ALL TABLES!")
		fmt.Print("Are you sure? (yes/no): ")

		var confirm string
		_, err = fmt.Scanln(&confirm)
		if err != nil {
			return
		}

		if confirm == "yes" {
			if err = migrations.Rollback(db); err != nil {
				fmt.Printf("❌ Rollback failed: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("❌ Rollback cancelled")
		}

	case "migrate:status":
		// Cek status migrasi
		if err = migrations.CheckMigrationStatus(db); err != nil {
			fmt.Printf("❌ Status check failed: %v\n", err)
			os.Exit(1)
		}

	case "db:seed":
		// Seed db
		if err = migrations.Seed(db); err != nil {
			fmt.Printf("❌ Seeding failed: %v\n", err)
			os.Exit(1)
		}

	case "migrate:fresh-seed":
		// Fresh migration + seed
		fmt.Println("⚠️  WARNING: This will DROP ALL TABLES and data, then seed!")
		fmt.Print("Are you sure? (yes/no): ")

		var confirm string
		_, err = fmt.Scanln(&confirm)
		if err != nil {
			return
		}

		if confirm == "yes" {
			if err = migrations.FreshSeed(db); err != nil {
				fmt.Printf("❌ Fresh seed failed: %v\n", err)
				os.Exit(1)
			}
		} else {
			fmt.Println("❌ Operation cancelled")
		}

	default:
		fmt.Printf("❌ Unknown command: %s\n\n", command)
		printUsage()
		os.Exit(1)
	}
}

func printUsage() {
	fmt.Println("Migration Tool - Database Migration Helper")
	fmt.Println("\nUsage:")
	fmt.Println("  go run cmd/migrate/main.go <command>")
	fmt.Println("\nAvailable commands:")
	fmt.Println("  migrate              - Run migrations (create/update tables)")
	fmt.Println("  migrate:fresh        - Drop all tables and re-run migrations")
	fmt.Println("  migrate:rollback     - Drop all tables")
	fmt.Println("  migrate:status       - Check migration status")
	fmt.Println("  db:seed              - Seed database with dummy data")
	fmt.Println("  migrate:fresh-seed   - Fresh migrate + seed")
	fmt.Println("\nExamples:")
	fmt.Println("  go run cmd/migrate/main.go migrate")
	fmt.Println("  go run cmd/migrate/main.go migrate:fresh")
	fmt.Println("  go run cmd/migrate/main.go db:seed")
}
