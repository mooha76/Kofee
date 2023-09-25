package db

import (
	"fmt"

	"github.com/mooha76/GoUser_Service-Grpc/config"
	"github.com/mooha76/GoUser_Service-Grpc/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDatabase(cfg *config.Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(
		model.User{},
	)

	if err != nil {
		return nil, err
	}
	return db, nil
}
