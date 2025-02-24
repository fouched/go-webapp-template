package handlers

import (
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"net/http"
	"strings"
)

func (h *Handlers) Search(w http.ResponseWriter, r *http.Request) {
	page, ok := h.App.Session.Get(r.Context(), "page").(string)
	// this should never happen...
	if !ok {
		h.App.ErrorLog.Println("No session data for search exiting ")
		return
	}

	pageNum := 0
	filter := strings.TrimLeft(r.URL.Query().Get("filter"), " ")
	data := make(map[string]interface{})

	if page == "customer" {
		customers, err := services.CustomerService(h.App).GetCustomerGrid(pageNum, filter)
		if err != nil {
			h.App.ErrorLog.Print(err)
			h.App.Session.Put(r.Context(), "error", "Search error")
		}
		data["customers"] = customers
	}

	intMap := make(map[string]int)
	intMap["pageNum"] = pageNum
	stringMap := make(map[string]string)
	stringMap["filter"] = filter
	stringMap["page"] = page
	_ = render.Partial(w, r, page+"-search", &render.TemplateData{
		Data:      data,
		IntMap:    intMap,
		StringMap: stringMap,
	}, page+"-row")
}
