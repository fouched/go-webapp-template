package repo

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
)

type testTeamRepo struct {
	App *config.App
}

func NewTestTeamRepo(a *config.App) TeamRepo {
	return &testTeamRepo{App: a}
}

func (r testTeamRepo) Create(t *models.Team) error {
	return nil
}
