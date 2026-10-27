package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	// Load .env if available.
	// Existing system environment variables are also supported.
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file loaded; using environment variables")
	}

	required := []string{
		"DB_USER",
		"DB_HOST",
		"DB_PORT",
		"DB_NAME",
	}

	for _, key := range required {
		if strings.TrimSpace(os.Getenv(key)) == "" {
			return fmt.Errorf("missing required environment variable: %s", key)
		}
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to open database: %w", err)
	}

	sqlDB, err := d.DB()
	if err != nil {
		return fmt.Errorf("failed to get SQL database: %w", err)
	}

	if err := sqlDB.Ping(); err != nil {
		sqlDB.Close()
		return fmt.Errorf("failed to connect to MySQL: %w", err)
	}

	db = d
	return nil
}

func GetDB() *gorm.DB {
	return db
}
