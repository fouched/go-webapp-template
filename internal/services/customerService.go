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
		customerService = &customerServicer{
			Repo: repo.NewCustomerRepo(a),
			App:  a,
		}
	}

	return customerService
}

func (s *customerServicer) GetCustomerGrid(page int, filter string) (*[]models.Customer, error) {
	customers, err := s.Repo.SelectCustomerGrid(page, filter)
	if err != nil {
		return nil, err
	}

	return customers, nil
}

func (s *customerServicer) GetCustomerById(id int64) (*models.Customer, error) {
	customer, err := s.Repo.SelectCustomerById(id)
	if err != nil {
		return nil, err
	}

	return customer, nil
}

func (s *customerServicer) CustomerInsert(customer *models.Customer) (int64, error) {
	id, err := s.Repo.CustomerInsert(customer)
	return id, err
}

func (s *customerServicer) CustomerUpdate(customer *models.Customer) error {
	return s.Repo.CustomerUpdate(customer)
}

func (s *customerServicer) DeleteCustomerById(id int64) error {
	return s.Repo.CustomerDelete(id)
}
