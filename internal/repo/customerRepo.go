package repo

import (
	"database/sql"
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
)

type postgresCustomerRepo struct {
	App *config.App
	DB  *sql.DB
}

func NewCustomerRepo(a *config.App) CustomerRepo {
	return &postgresCustomerRepo{
		App: a,
		DB:  a.DB.SQL,
	}
}

func (r *postgresCustomerRepo) Create(c *models.Customer) error {
	r.App.InfoLog.Println("in postgresCustomerRepo.Create")

	//ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	//defer cancel()
	//
	//query := ``
	//rows, err := r.DB.QueryContext(ctx, query)

	return nil
}
