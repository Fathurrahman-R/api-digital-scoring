package migrations

import (
	"api-digital-scoring/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

// GetAllModels mengembalikan semua model yang perlu dimigrasikan
func GetAllModels() []interface{} {
	return []interface{}{
		&entity.User{},
		&entity.RefreshToken{},
		// Tambahkan model lain di sini sesuai kebutuhan
		// &entity.Product{},
		// &entity.Category{},
	}
}

// Migrate melakukan migrasi biasa (auto migrate)
// Fungsi ini akan membuat tabel baru atau menambahkan kolom baru
// tanpa menghapus data yang sudah ada
func Migrate(db *gorm.DB) error {
	fmt.Println("🚀 Starting database migration...")

	models := GetAllModels()

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %w", model, err)
		}
		fmt.Printf("✅ Migrated: %T\n", model)
	}

	fmt.Println("✨ Migration completed successfully!")
	return nil
}

// FreshMigrate melakukan fresh migration (drop all tables dan recreate)
// HATI-HATI: Fungsi ini akan menghapus semua data!
func FreshMigrate(db *gorm.DB) error {
	fmt.Println("⚠️  Starting FRESH migration (dropping all tables)...")

	models := GetAllModels()

	// Drop semua tabel
	fmt.Println("🗑️  Dropping all tables...")
	if err := db.Migrator().DropTable(models...); err != nil {
		return fmt.Errorf("failed to drop tables: %w", err)
	}

	// Recreate semua tabel
	fmt.Println("🔨 Creating tables...")
	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %w", model, err)
		}
		fmt.Printf("✅ Created: %T\n", model)
	}

	fmt.Println("✨ Fresh migration completed successfully!")
	return nil
}

// Rollback menghapus semua tabel
func Rollback(db *gorm.DB) error {
	fmt.Println("🔄 Rolling back migrations...")

	models := GetAllModels()

	if err := db.Migrator().DropTable(models...); err != nil {
		return fmt.Errorf("failed to rollback: %w", err)
	}

	fmt.Println("✅ Rollback completed successfully!")
	return nil
}

// MigrateSpecific melakukan migrasi untuk model tertentu saja
func MigrateSpecific(db *gorm.DB, models ...interface{}) error {
	fmt.Println("🚀 Starting specific migration...")

	for _, model := range models {
		if err := db.AutoMigrate(model); err != nil {
			return fmt.Errorf("failed to migrate %T: %w", model, err)
		}
		fmt.Printf("✅ Migrated: %T\n", model)
	}

	fmt.Println("✨ Specific migration completed!")
	return nil
}

// CheckMigrationStatus mengecek status tabel di database
func CheckMigrationStatus(db *gorm.DB) error {
	fmt.Println("📊 Checking migration status...")

	models := GetAllModels()

	for _, model := range models {
		_ = db.Migrator().CurrentDatabase()
		hasTable := db.Migrator().HasTable(model)

		if hasTable {
			fmt.Printf("✅ Table exists: %T\n", model)
		} else {
			fmt.Printf("❌ Table missing: %T\n", model)
		}
	}

	return nil
}
