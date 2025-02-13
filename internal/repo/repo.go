package repo

import (
	"github.com/fouched/go-webapp-template/internal/models"
	"time"
)

// DbTimeout should be lowered in production - 30 secs max
const DbTimeout = time.Minute * 2
const PageSize = 20

type CustomerRepo interface {
	Create(customer *models.Customer) error
	SelectCustomerGrid(page int) (*[]models.Customer, error)
	SelectCustomerGridWithFilter(page int, filter string) (*[]models.Customer, error)
	SelectCustomerById(int int64) (*models.Customer, error)
	UpdateCustomer(customer *models.Customer) error
}
