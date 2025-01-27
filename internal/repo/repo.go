package repo

import (
	"github.com/fouched/go-webapp-template/internal/models"
	"time"
)

// DbTimeout should be lowered in production - 30 secs max
const DbTimeout = time.Minute * 2
const PageSize = 20

type CustomerRepo interface {
	Create(customer *models.Customer) error
}

type TeamRepo interface {
	Create(team *models.Team) error
}
