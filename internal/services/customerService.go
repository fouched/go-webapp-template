package services

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/repo"
)

var customerService *customerServicer

type customerServicer struct {
	Repo repo.CustomerRepo
	App  *config.App
}

func CustomerService(a *config.App) CustomerServicer {
	if customerService == nil {
		a.InfoLog.Println("Creating customer service")
		customerService = &customerServicer{
			Repo: repo.NewCustomerRepo(a),
			App:  a,
		}
	}

	return customerService
}

func (s *customerServicer) CreateCustomer(c *models.Customer) error {
	err := s.Repo.Create(c)
	if err != nil {
		s.App.ErrorLog.Print(err)
		return err
	}

	return nil
}

func (s *customerServicer) GetCustomerGrid(page int) (*[]models.Customer, error) {
	customers, err := s.Repo.SelectCustomerGrid(page)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *customerServicer) GetCustomerGridWithFilter(page int, filter string) (*[]models.Customer, error) {
	customers, err := s.Repo.SelectCustomerGridWithFilter(page, filter)
	if err != nil {
		return nil, err
	}

	return customers, nil
}
