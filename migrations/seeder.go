package migrations

import (
	"api-digital-scoring/internal/entity"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Seed mengisi database dengan data dummy untuk testing
func Seed(db *gorm.DB) error {
	fmt.Println("🌱 Starting database seeding...")

	// Seed Users
	if err := seedUsers(db); err != nil {
		return err
	}

	// Tambahkan seeder lain di sini
	// if err := seedProducts(db); err != nil {
	//     return err
	// }

	fmt.Println("✨ Seeding completed successfully!")
	return nil
}

// seedUsers mengisi tabel users dengan data dummy
func seedUsers(db *gorm.DB) error {
	fmt.Println("👥 Seeding users...")
	hashed, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	users := []entity.User{
		{
			Fullname: "Admin",
			Username: "admin",
			Email:    "admin@example.com",
			Password: string(hashed), // Hash password yang sebenarnya
		},
	}

	for _, user := range users {
		// Check apakah user sudah ada
		var existingUser entity.User
		result := db.Where("email = ?", user.Email).First(&existingUser)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			// User belum ada, create
			if err := db.Create(&user).Error; err != nil {
				return fmt.Errorf("failed to seed user %s: %w", user.Email, err)
			}
			fmt.Printf("✅ Created user: %s\n", user.Email)
		} else {
			fmt.Printf("⏭️  User already exists: %s\n", user.Email)
		}
	}

	return nil
}

// FreshSeed melakukan fresh migration dan seeding
func FreshSeed(db *gorm.DB) error {
	// Drop dan recreate tables
	if err := FreshMigrate(db); err != nil {
		return err
	}

	// Seed data
	if err := Seed(db); err != nil {
		return err
	}

	return nil
}
