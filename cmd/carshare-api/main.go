package main

import (
	"carshare-api/internal/config"
	"fmt"
	"log"
	"os"
)

func main() {

	err := os.Setenv("CONFIG_PATH", "./config/local.yaml")
	if err != nil {
		log.Fatal("Can not set env var")
	}

	// init config
	cfg := config.MustLoad()
	fmt.Println(cfg.PostgresDSN)
	// init storage

	// init router

}
