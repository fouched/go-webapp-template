package handlers

import (
	"fmt"
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func (h *Handlers) CustomerGrid(w http.ResponseWriter, r *http.Request) {
	h.App.Session.Put(r.Context(), "searchTarget", "customer-search")

	h.App.InfoLog.Println("loading customer grid")

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
		Data:      data,
		IntMap:    intMap,
		StringMap: stringMap,
	})
}

func (h *Handlers) CustomerDetails(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	customer, err := services.CustomerService(h.App).GetCustomerById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error getting customer")
	}

	data := make(map[string]interface{})
	data["Customer"] = customer

	_ = render.Partial(w, r, "customer-details", &render.TemplateData{
		Data: data,
	})
}

func (h *Handlers) CustomerAddGet(w http.ResponseWriter, r *http.Request) {
	_ = render.Template(w, r, "customer-add", &render.TemplateData{})
}

func (h *Handlers) CustomerAddPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error parsing form for customer insert")
	}

	customer := models.Customer{
		CustomerName: r.Form.Get("customerName"),
		Tel:          r.Form.Get("tel"),
		Email:        r.Form.Get("email"),
		Address1:     r.Form.Get("address1"),
		Address2:     r.Form.Get("address2"),
		Address3:     r.Form.Get("address3"),
		PostCode:     r.Form.Get("postCode"),
		UpdatedAt:    time.Now(),
	}

	id, err := services.CustomerService(h.App).CustomerInsert(&customer)
	if err != nil {
		h.App.ErrorLog.Print(err)
		if strings.Contains(err.Error(), "duplicate") {
			h.App.Session.Put(r.Context(), "error", "Error customer already exists")
		} else {
			h.App.Session.Put(r.Context(), "error", "Error inserting customer")
		}
	} else {
		// in a real app we might to do something with an inserted record
		h.App.InfoLog.Println(fmt.Sprintf("Inserted customer with id: %d", id))
		h.App.Session.Put(r.Context(), "success", "Successfully added customer")
		customer.ID = id
	}

	http.Redirect(w, r, "/customers", http.StatusSeeOther)
}

func (h *Handlers) CustomerUpdate(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error updating customer")
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
		h.App.Session.Put(r.Context(), "error", "Error updating customer")
	}

	data := make(map[string]interface{})
	data["Customer"] = customer

	h.App.Session.Put(r.Context(), "success", "Customer updated successfully")
	_ = render.Partial(w, r, "customer-row", &render.TemplateData{
		Data: data,
	})

}

func (h *Handlers) CustomerDelete(w http.ResponseWriter, r *http.Request) {

	id, _ := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	err := services.CustomerService(h.App).DeleteCustomerById(id)
	if err != nil {
		h.App.ErrorLog.Print(err)
		h.App.Session.Put(r.Context(), "error", "Error deleting customer")

		// load the customer again because we normally remove the row from the UI
		customer, _ := services.CustomerService(h.App).GetCustomerById(id)
		data := make(map[string]interface{})
		data["Customer"] = customer
		_ = render.Partial(w, r, "customer-row", &render.TemplateData{
			Data: data,
		})
	} else {
		h.App.Session.Put(r.Context(), "success", "Customer deleted successfully")
		_ = render.Partial(w, r, "customer-delete", &render.TemplateData{})
	}

}
