package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"net/http"
	"strconv"
)

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	t, ok := h.App.Session.Get(r.Context(), "searchTarget").(string)
	// this should never happen...
	if !ok {
		h.App.ErrorLog.Println("No session data for search exiting ")
		return
	}

	page := 0
	p := r.URL.Query().Get("page")
	if p != "" {
		page, _ = strconv.Atoi(p)
	}

	q := r.URL.Query().Get("q")
	data := make(map[string]interface{})

	if t == "customer-search" {
		customers, err := services.CustomerService(h.App).GetCustomerGridWithFilter(page, q)
		if err != nil {
			h.App.ErrorLog.Print(err)
			return
		}
		data["customers"] = customers
	}

	intMap := make(map[string]int)
	intMap["page"] = page
	stringMap := make(map[string]string)
	stringMap["q"] = q
	_ = render.Partial(w, r, t, &render.TemplateData{
		Data:      data,
		IntMap:    intMap,
		StringMap: stringMap,
	})
}
