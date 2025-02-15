package handlers

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/render"
	"net/http"
)

// Instance the setup for all handlers
var Instance *Handlers

type Handlers struct {
	App *config.App
}

// NewHandlerConfig set the handler config and services
func NewHandlerConfig(a *config.App) *Handlers {
	return &Handlers{
		App: a,
	}
}

// NewHandlers set
func NewHandlers(h *Handlers) {
	Instance = h
}

func (h *Handlers) Toast(w http.ResponseWriter, r *http.Request) {
	_ = render.Partial(w, r, "toast", &render.TemplateData{})
}
