package main

import (
	"github.com/fouched/go-webapp-template/internal/handlers"
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net/http"
	"time"
)

func routes() http.Handler {
	r := chi.NewRouter()
	addMiddleware(r)

	// routes
	r.Get("/", handlers.Instance.Home)
	r.Get("/search", handlers.Instance.Search)

	r.Route("/customer", func(r chi.Router) {
		r.Get("/", handlers.Instance.CustomerGrid)
		r.Get("/{id}", handlers.Instance.CustomerDetails)
		r.Get("/add", handlers.Instance.CustomerAddGet)
		r.Post("/add", handlers.Instance.CustomerAddPost)
		r.Post("/{id}/update", handlers.Instance.CustomerUpdate)
		r.Delete("/{id}", handlers.Instance.CustomerDelete)
	})

	// add ability to render static resources
	r.Handle("/static/*", http.FileServer(http.FS(render.EmbedFS())))

	// to externalize above rather use below, and move static files to top level directory of project
	//fileServer := http.FileServer(http.Dir("./static/"))
	//r.Handle("/static/*", http.StripPrefix("/static", fileServer))

	return r
}

func addMiddleware(r *chi.Mux) {
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
}
