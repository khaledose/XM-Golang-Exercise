package db

import (
	"xm/companies/config"
	"xm/companies/internal/models"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabaseConnection(logger *zap.SugaredLogger, config config.DatabaseConfigurations) *gorm.DB {
	db, err := gorm.Open(postgres.Open(config.Dsn), &gorm.Config{})
	if err != nil {
		logger.Fatalf("Failed to open connection to database, with error: %s", err)
		panic(err)
	}

	err = db.AutoMigrate(&models.Company{})
	if err != nil {
		logger.Fatalf("Failed to migrate models in the database, with error: %s", err)
		panic(err)
	}

	return db
}
