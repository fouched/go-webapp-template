package main

import (
	"github.com/fouched/go-webapp-template/internal/handlers"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func routes() http.Handler {
	r := chi.NewRouter()

	r.Get("/", handlers.Instance.Home)

	return r
}
