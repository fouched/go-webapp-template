package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (h *Handlers) CustomerGrid(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "searchTarget", "customer-rows")

	page := 0
	p := r.URL.Query().Get("page")

	if p != "" {
		page, _ = strconv.Atoi(p)
	}

	customers, err := services.CustomerService(h.App).GetCustomerGrid(page)
	if err != nil {
		h.App.ErrorLog.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["customers"] = customers
	intMap := make(map[string]int)
	intMap["page"] = page

	_ = render.Template(w, r, "customers", &render.TemplateData{
		Data:   data,
		IntMap: intMap,
	}) //todo add partial for HTMX customer rows
}

func (h *Handlers) CustomerDetails(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	h.App.InfoLog.Printf("select customer: %d\n", id)

	_ = render.Partial(w, r, "customer-details", &render.TemplateData{})
}
