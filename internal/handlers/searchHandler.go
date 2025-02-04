package handlers

import (
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

	h.App.InfoLog.Println("Filter:", q, " Page:", page, "Template:", t)
}
