package repo

import "github.com/fouched/go-webapp-template/internal/models"

type CustomerRepo interface {
	Create(customer *models.Customer) error
}

type TeamRepo interface {
	Create(team *models.Team) error
}
