package repo

import (
	"database/sql"
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
)

type postgresCustomerRepo struct {
	DB  *sql.DB
	App *config.App
}

func NewCustomerRepo(a *config.App, db *sql.DB) CustomerRepo {
	return &postgresCustomerRepo{
		DB:  db,
		App: a,
	}
}

func (r *postgresCustomerRepo) Create(c *models.Customer) error {
	r.App.InfoLog.Println("in postgresCustomerRepo.Create")

	//r.DB.Query()

	return nil
}
