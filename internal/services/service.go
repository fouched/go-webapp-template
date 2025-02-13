package services

import "github.com/fouched/go-webapp-template/internal/models"

type CustomerServicer interface {
	CreateCustomer(request *models.Customer) error
	GetCustomerGrid(page int) (*[]models.Customer, error)
	GetCustomerGridWithFilter(page int, filter string) (*[]models.Customer, error)
	GetCustomerById(id int64) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
}
