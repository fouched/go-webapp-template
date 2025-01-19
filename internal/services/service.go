package services

import "github.com/fouched/go-webapp-template/internal/models"

type CustomerServicer interface {
	CreateCustomer(request *models.Customer) error
}

type TeamServicer interface {
	CreateTeam(request *models.Team) error
}
