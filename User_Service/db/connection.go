package db

import (
	_ "database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/mooha76/Kofee/User_Service/config"
	"github.com/mooha76/Kofee/User_Service/model"
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

func InitializeSQLXDatabase(cfg *config.Config) (*sqlx.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s sslmode=%s",
		cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword, cfg.Sslmode)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("error opening database: %v", err)
	}

	return db, nil
}
