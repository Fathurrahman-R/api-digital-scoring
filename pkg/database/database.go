package database

import (
	"api-digital-scoring/internal/config"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewMySQL(cfg config.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User, cfg.Password, cfg.Host, cfg.Port, cfg.DBName)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		return nil, fmt.Errorf("failed connected to database: %w", err)
	}

	return db, nil
}
