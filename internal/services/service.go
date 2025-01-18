package services

import "github.com/fouched/go-webapp-template/internal/models"

type CustomerService interface {
	CreateCustomer(request *models.Customer) error
}

type TeamService interface {
	CreateTeam(request *models.Team) error
}
