package services

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
	"github.com/fouched/go-webapp-template/internal/repo"
)

var teamService *teamServicer

type teamServicer struct {
	Repo repo.TeamRepo
	App  *config.App
}

type CreateTeamRequest struct {
	ID int64
}

func TeamService(a *config.App) TeamServicer {
	if teamService == nil {
		a.InfoLog.Println("Creating team service")
		teamService = &teamServicer{
			Repo: repo.NewTeamRepo(a),
			App:  a,
		}
	} else {
		a.InfoLog.Println("RE-USE team service")
	}
	return teamService
}

func (s *teamServicer) CreateTeam(t *models.Team) error {
	s.App.InfoLog.Println("in services.CreateTeam")

	err := s.Repo.Create(t)
	if err != nil {
		s.App.ErrorLog.Print(err)
		return err
	}
	return nil
}
