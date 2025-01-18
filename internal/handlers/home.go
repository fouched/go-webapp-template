package handlers

import (
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/render"
	"net/http"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {

	h.InfoLog.Println("FooBar")
	err := h.CustomerService.CreateCustomer(&models.Customer{})
	if err != nil {
		return
	}

	err = h.TeamService.CreateTeam(&models.Team{})
	if err != nil {
		return
	}

	render.Template(w, r, "home", &render.TemplateData{})
}
