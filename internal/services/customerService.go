package services

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/driver"
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/repo"
)

type customerServiceInstance struct {
	Repo repo.CustomerRepo
	App  *config.App
}

func NewCustomerService(a *config.App, db *driver.DB) CustomerService {
	return &customerServiceInstance{
		Repo: repo.NewCustomerRepo(a, db.SQL),
		App:  a,
	}
}

func (s *customerServiceInstance) CreateCustomer(c *models.Customer) error {
	s.App.InfoLog.Println("in services.CreateCustomer")

	err := s.Repo.Create(c)
	if err != nil {
		s.App.ErrorLog.Print(err)
		return err
	}

	return nil
}
