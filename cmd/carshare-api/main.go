package main

import (
	"carshare-api/internal/config"
	"carshare-api/internal/storage/pgsql"
	"log"
	"os"
)

func main() {

	// init config
	err := os.Setenv("CONFIG_PATH", "./config/local.yaml")
	if err != nil {
		log.Fatal("Can not set env var")
	}
	cfg := config.MustLoad()

	// init storage
	storage, err := pgsql.New(cfg.PostgresDSN)
	if err != nil {
		log.Fatal("Error on storage init\n", err)
	}
	_ = storage

	// init router

}
