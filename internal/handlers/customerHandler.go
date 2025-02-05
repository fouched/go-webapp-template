package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
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
