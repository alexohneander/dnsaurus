package database

import (
	"fmt"

	"github.com/alexohneander/dnsaurus/config"
	"github.com/alexohneander/dnsaurus/model"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDB connect to db
func ConnectDB() {
	var err error

	path := config.Config("DB_PATH")
	if path == "" {
		panic("failed to parse database path")
	}

	DB, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	fmt.Println("Connection Opened to Database")
	DB.AutoMigrate(&model.User{})
	fmt.Println("Database Migrated")
}

func GetDB() (*gorm.DB, error) {
	var err error

	path := config.Config("DB_PATH")
	if path == "" {
		panic("failed to parse database path")
	}

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db, err
}
