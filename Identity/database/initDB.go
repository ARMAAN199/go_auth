package database

import (
	"fmt"

	"github.com/ARMAAN199/Go_EcomApi/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *config.Config) (*gorm.DB, error) {

	// var dbconfig = DatabaseConfig{}
	dsn := DbDSN(cfg)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("[error] failed to connect to database:", err)
		return nil, err
	}

	return db, nil
}
