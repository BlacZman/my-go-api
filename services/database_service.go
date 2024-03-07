package services

import (
	"my-go-api/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService struct {
	db     *gorm.DB
	config AppConfigService
}

func NewDatabaseConnection(config AppConfigService) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.DSN()), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func NewDatabaseService(db *gorm.DB, config AppConfigService) DatabaseService {
	// Migration
	db.AutoMigrate(&models.User{})

	return DatabaseService{
		db:     db,
		config: config,
	}
}
