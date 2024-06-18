package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ARMAAN199/Go_EcomApi/config"
	"github.com/ARMAAN199/Go_EcomApi/database"
	router "github.com/ARMAAN199/Go_EcomApi/routers"
)

func main() {
	fmt.Println("Before Server Start")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = ":8088"
	}

	cfg := config.InitConfig()

	db, err := database.InitDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = http.ListenAndServe(PORT, *router.ReturnRouter(db))
	if err != nil {
		log.Fatal(err)
	}
}
