package database

import (
	"github.com/vitaly06/geo-weather-bot/internal/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode((logger.Silent)),
	})

	return db, err
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(&domain.User{})
}
