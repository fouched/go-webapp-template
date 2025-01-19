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
	} else {
		a.InfoLog.Println("RE-USE customer service")
	}
	return customerService
}

func (s *customerServicer) CreateCustomer(c *models.Customer) error {
	s.App.InfoLog.Println("in services.CreateCustomer")

	err := s.Repo.Create(c)
	if err != nil {
		s.App.ErrorLog.Print(err)
		return err
	}

	return nil
}
