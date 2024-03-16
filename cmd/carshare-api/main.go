package main

import (
	"carshare-api/internal/config"
	"carshare-api/internal/http-server/handlers/getCarLocations"
	"carshare-api/internal/http-server/handlers/getCars"
	"carshare-api/internal/http-server/handlers/postCarLocation"
	"carshare-api/internal/storage/pgsql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
	"os"
)

func main() {

	// Config setup
	err := os.Setenv("CONFIG_PATH", "./config/local.yaml")
	if err != nil {
		log.Fatal("Can not set env var")
	}
	cfg := config.MustLoad()

	// Storage setup
	storage, err := pgsql.New(cfg.PostgresDSN)
	if err != nil {
		log.Fatal("Error on storage init\n", err)
	}

	// Router setup
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	// Dummy route
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is ready for battle ;)"))
	})

	// Main routes
	router.Get("/locations", getCarLocations.New(storage))
	router.Get("/cars", getCars.New(storage))
	router.Post("/postLocation", postCarLocation.New(storage))

	log.Println("Starting server at", cfg.Address)

	srv := &http.Server{
		Addr:         cfg.Address,
		Handler:      router,
		ReadTimeout:  cfg.HTTPServer.Timeout,
		WriteTimeout: cfg.HTTPServer.Timeout,
		IdleTimeout:  cfg.HTTPServer.IdleTimeout,
	}

	srv.ListenAndServe()

}
