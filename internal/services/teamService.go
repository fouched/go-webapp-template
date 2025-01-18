package services

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/driver"
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/repo"
)

type teamServiceInstance struct {
	Repo repo.TeamRepo
	App  *config.App
}

type CreateTeamRequest struct {
	ID int64
}

func NewTeamService(a *config.App, db *driver.DB) TeamService {
	return &teamServiceInstance{
		Repo: repo.NewTeamRepo(a, db.SQL),
		App:  a,
	}
}

func (s *teamServiceInstance) CreateTeam(t *models.Team) error {
	s.App.InfoLog.Println("in services.CreateTeam")

	return s.Repo.Create(t)
}
