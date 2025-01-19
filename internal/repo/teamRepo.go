package repo

import (
	"database/sql"
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
)

type postgresTeamRepo struct {
	App *config.App
	DB  *sql.DB
}

func NewTeamRepo(a *config.App) TeamRepo {
	return &postgresTeamRepo{
		App: a,
		DB:  a.DB.SQL,
	}
}

func (r *postgresTeamRepo) Create(t *models.Team) error {
	r.App.InfoLog.Println("in postgresCustomerRepo.Create")

	return nil
}
