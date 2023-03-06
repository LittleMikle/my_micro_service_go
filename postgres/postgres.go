package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectToPostgres() (db *gorm.DB, err error) {
	dbURL := "postgres://pg:123@localhost:5432/crud"

	db, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed with gorm.Open: %w", err)
	}

	err = db.AutoMigrate()
	if err != nil {
		return nil, fmt.Errorf("failed with AutoMigrate: %w", err)
	}
	return db, nil
}
