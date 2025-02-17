package services

import "github.com/fouched/go-webapp-template/internal/models"

type CustomerServicer interface {
	GetCustomerGrid(page int, filter string) (*[]models.Customer, error)
	GetCustomerById(id int64) (*models.Customer, error)
	CustomerInsert(customer *models.Customer) (int64, error)
	CustomerUpdate(customer *models.Customer) error
}
