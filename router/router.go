package router

import (
	"app/config"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func AppRouter() http.Handler {
	app := chi.NewRouter()

	// A good base middleware stack
	app.Use(middleware.RequestID)
	app.Use(middleware.RealIP)
	app.Use(middleware.Logger)
	app.Use(middleware.Recoverer)

	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	app.Route("/api/v1", apiV1)

	log.Printf(
		"Server h-learning-upload-mp4 starting success! URL: http://%s:%s",
		config.GetAppHost(),
		config.GetAppPort(),
	)

	return app
}
