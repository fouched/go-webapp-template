package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"net/http"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "searchTarget", "home")

	_ = render.Template(w, r, ".", "home", &render.TemplateData{})
}
