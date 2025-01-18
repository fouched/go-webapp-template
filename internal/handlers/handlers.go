package handlers

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/driver"
	"github.com/fouched/go-webapp-template/internal/services"
	"log"
)

// Instance the setup for all handlers
var Instance *Handlers

type Handlers struct {
	InfoLog         *log.Logger
	ErrorLog        *log.Logger
	CustomerService services.CustomerService
	TeamService     services.TeamService
}

// NewHandlerConfig set the handler config and services
func NewHandlerConfig(c *config.App, db *driver.DB) *Handlers {
	return &Handlers{
		InfoLog:         c.InfoLog,
		ErrorLog:        c.ErrorLog,
		CustomerService: services.NewCustomerService(c, db),
		TeamService:     services.NewTeamService(c, db),
	}
}

// NewHandlers set
func NewHandlers(h *Handlers) {
	Instance = h
}
