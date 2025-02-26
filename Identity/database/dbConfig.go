package database

import (
	"fmt"

	"github.com/ARMAAN199/Go_EcomApi/config"
)

// type DatabaseConfig struct {
// 	Name      string
// 	Host      string
// 	Port      int
// 	Username  string
// 	Password  string
// 	Driver    string
// 	EnableLog bool
// }

// DbDSN takes DB configuration and returns a DB connection string
func DbDSN(cfg *config.Config) string {
	dbConfigString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", cfg.DBHost, cfg.DBUser, cfg.DBPassword, cfg.DBname, cfg.DBPort, cfg.SSLMode)
	return dbConfigString
}
