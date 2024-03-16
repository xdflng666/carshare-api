package main

import (
	"carshare-api/internal/config"
	"carshare-api/internal/http-server/handlers/getCarLocations"
	"carshare-api/internal/storage/pgsql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"log"
	"net/http"
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

	// init router
	router := chi.NewRouter()
	router.Use(middleware.URLFormat)
	router.Use(middleware.Logger)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("API is ready for battle ;)"))
	})
	router.Get("/locations", getCarLocations.New(storage))

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
