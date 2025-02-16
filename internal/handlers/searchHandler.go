package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"net/http"
	"strings"
)

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	t, ok := h.App.Session.Get(r.Context(), "searchTarget").(string)
	// this should never happen...
	if !ok {
		h.App.ErrorLog.Println("No session data for search exiting ")
		return
	}

	page := 0
	filter := strings.TrimLeft(r.URL.Query().Get("filter"), " ")
	data := make(map[string]interface{})

	if t == "customer-search" {
		customers, err := services.CustomerService(h.App).GetCustomerGrid(page, filter)
		if err != nil {
			h.App.ErrorLog.Print(err)
			h.App.Session.Put(r.Context(), "error", "Search error")
		}
		data["customers"] = customers
	}

	intMap := make(map[string]int)
	intMap["page"] = page
	stringMap := make(map[string]string)
	stringMap["filter"] = filter
	_ = render.Partial(w, r, t, &render.TemplateData{
		Data:      data,
		IntMap:    intMap,
		StringMap: stringMap,
	})
}
