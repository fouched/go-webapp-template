package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"net/http"
)

func (h *Handlers) CustomerGrid(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "searchTarget", "customer-rows")

	data := make(map[string]interface{})
	intMap := make(map[string]int)
	_ = render.Template(w, r, "customers", &render.TemplateData{
		Data:   data,
		IntMap: intMap,
	}) //todo add partial for HTMX customer rows
}
