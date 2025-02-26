package main

import (
	"log"
	"os"

	"github.com/ARMAAN199/Go_EcomApi/config"
	"github.com/ARMAAN199/Go_EcomApi/database"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {

	cfg := config.InitConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	pgDB, err := db.DB()
	if err != nil {
		log.Fatalf("failed to get db: %v", err)
	}

	driver, err := postgres.WithInstance(pgDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("could not start sql migration: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrate/migrations",
		"postgres",
		driver,
	)

	if err != nil {
		log.Fatalf("migration failed before up or down: %v", err)
	}

	args := os.Args
	userArg := args[len(args)-1]
	if userArg == "up" {
		if err := m.Up(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("an error occurred while syncing the database: %v", err)
		}
		log.Println("Migrations applied successfully!")
	} else if userArg == "down" {
		if err := m.Down(); err != nil && err != migrate.ErrNoChange {
			log.Fatalf("an error occurred while syncing the database: %v", err)
		}
		log.Println("Migrations applied successfully!")
	}

}
