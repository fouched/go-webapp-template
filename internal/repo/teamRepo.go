package repo

import (
	"database/sql"
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
)

type postgresTeamRepo struct {
	DB  *sql.DB
	App *config.App
}

func NewTeamRepo(a *config.App, db *sql.DB) TeamRepo {
	return &postgresTeamRepo{
		DB:  db,
		App: a,
	}
}

func (r *postgresTeamRepo) Create(t *models.Team) error {
	r.App.InfoLog.Println("in postgresCustomerRepo.Create")

	return nil
}
