package main

import (
	"carshare-api/internal/config"
	"carshare-api/internal/http-server/handlers/getCarLocations"
	"carshare-api/internal/http-server/handlers/getCars"
	"carshare-api/internal/http-server/handlers/getStory"
	"carshare-api/internal/http-server/handlers/postCarLocation"
	"carshare-api/internal/storage/pgsql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
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
	// Basic CORS
	// for more ideas, see: https://developer.github.com/v3/#cross-origin-resource-sharing
	router.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Protected routes
	router.Route("/api", func(r chi.Router) {

		// Define the basic auth middleware
		r.Use(middleware.BasicAuth("carshare-api", storage.GetAuthCredentials()))
		r.Get("/", func(writer http.ResponseWriter, request *http.Request) {
			writer.Write([]byte("Successfully got response from protected route root"))
		})

		// Main routes
		r.Get("/locations", getCarLocations.New(storage))
		r.Get("/cars", getCars.New(storage))
		r.Get("/story/{car_uuid}", getStory.New(storage))
		r.Post("/postLocation", postCarLocation.New(storage))
	})

	// Ping alive
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("I'm still alive"))
	})

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
