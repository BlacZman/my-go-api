package services

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseService struct {
	db     *gorm.DB
	config AppConfigService
}

func NewDatabaseService(config AppConfigService) DatabaseService {
	db, err := gorm.Open(postgres.Open(config.dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	return DatabaseService{
		db:     db,
		config: config,
	}
}
