package main

import (
	"github.com/fouched/go-webapp-template/internal/handlers"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func routes() http.Handler {
	r := chi.NewRouter()

	// sessions
	r.Use(SessionLoad)

	// CORS
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins: []string{"https://*", "http://*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Recover from panics, logs the panic, and returns HTTP 500
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// add ability to render static resources
	fileServer := http.FileServer(http.Dir("./static/"))
	r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	r.Get("/", handlers.Instance.Home)
	r.Get("/toast", handlers.Instance.Toast)
	r.Get("/search", handlers.Instance.Search)
	r.Get("/customers", handlers.Instance.CustomerGrid)
	r.Get("/customers/{id}", handlers.Instance.CustomerDetails)
	r.Post("/customers/{id}/update", handlers.Instance.CustomerUpdate)

	return r
}
