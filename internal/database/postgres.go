package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"wb_tech_level_0/config"
	"wb_tech_level_0/pkg/model"
)

// Подключение к бд и автомиграции

func ConnectPostgres(cfg *config.Config) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(cfg.PostgresDSN), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&model.Order{}, &model.Item{}, &model.Delivery{}, &model.Payment{}); err != nil {
		return nil, err
	}

	return db, nil
}
