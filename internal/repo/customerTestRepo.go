package repo

import (
	"github.com/fouched/go-webapp-template/internal/config"
	"github.com/fouched/go-webapp-template/internal/models"
)

type testCustomerRepo struct {
	App *config.App
}

func NewTestCustomerRepo(a *config.App) CustomerRepo {
	return &testCustomerRepo{
		App: a,
	}
}

func (r *testCustomerRepo) SelectCustomerGrid(page int, filter string) (*[]models.Customer, error) {
	r.App.InfoLog.Println("in TEST customerRepo.SelectCustomerGrid")
	return nil, nil
}

func (r *testCustomerRepo) SelectCustomerById(id int64) (*models.Customer, error) {
	return nil, nil
}

func (r *testCustomerRepo) CustomerInsert(customer *models.Customer) (int64, error) {
	return 0, nil
}

func (r *testCustomerRepo) CustomerUpdate(customer *models.Customer) error {
	return nil
}
