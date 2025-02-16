package handlers

import (
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"time"
)

func (h *Handlers) CustomerGrid(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "searchTarget", "customer-search")

	page := 0
	p := r.URL.Query().Get("page")
	filter := r.URL.Query().Get("filter")

	if p != "" {
		page, _ = strconv.Atoi(p)
	}

	customers, err := services.CustomerService(h.App).GetCustomerGrid(page, filter)

	if err != nil {
		h.App.ErrorLog.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["customers"] = customers
	intMap := make(map[string]int)
	intMap["page"] = page
	stringMap := make(map[string]string)
	stringMap["filter"] = filter

	_ = render.Template(w, r, "customers", &render.TemplateData{
		Data:   data,
		IntMap: intMap,
	})
}

func (h *Handlers) CustomerDetails(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	customer, err := services.CustomerService(h.App).GetCustomerById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["Customer"] = customer

	_ = render.Partial(w, r, "customer-details", &render.TemplateData{
		Data: data,
	})
}

func (h *Handlers) CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Customer could not be updated")
		http.Redirect(w, r, "/customers", http.StatusSeeOther)
		return
	}

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	customer := models.Customer{
		ID:           id,
		CustomerName: r.Form.Get("customerName"),
		Tel:          r.Form.Get("tel"),
		Email:        r.Form.Get("email"),
		Address1:     r.Form.Get("address1"),
		Address2:     r.Form.Get("address2"),
		Address3:     r.Form.Get("address3"),
		PostCode:     r.Form.Get("postCode"),
		UpdatedAt:    time.Now(),
	}

	err = services.CustomerService(h.App).CustomerUpdate(&customer)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Customer could not be updated")
		http.Redirect(w, r, "/customers", http.StatusSeeOther)
		return
	}

	data := make(map[string]interface{})
	data["Customer"] = customer

	h.App.Session.Put(r.Context(), "success", "Customer updated successfully")
	_ = render.Partial(w, r, "customer-row", &render.TemplateData{
		Data: data,
	})

}
