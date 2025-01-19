package handlers

import (
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/render"
	"github.com/fouched/go-webapp-template/internal/services"
	"net/http"
)

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {

	h.App.InfoLog.Println("FooBar")

	err := services.CustomerService(h.App).CreateCustomer(&models.Customer{})
	if err != nil {
		return
	}

	ts := services.TeamService(h.App)
	err = ts.CreateTeam(&models.Team{})
	if err != nil {
		return
	}

	render.Template(w, r, "home", &render.TemplateData{})
}
