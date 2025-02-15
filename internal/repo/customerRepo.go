package repo

import (
	"context"
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

func (r *postgresCustomerRepo) SelectCustomerGrid(page int) (*[]models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	s := `
		select c.id, c.customer_name, c.tel, c.email 
		from customer c
		order by c.customer_name
		limit $1 offset $2
	`

	rows, err := r.DB.QueryContext(ctx, s, PageSize, page*PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		err := rows.Scan(
			&c.ID,
			&c.CustomerName,
			&c.Tel,
			&c.Email,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	return &customers, nil
}

func (r *postgresCustomerRepo) SelectCustomerGridWithFilter(page int, filter string) (*[]models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	s := `
		select c.id, c.customer_name, c.tel, c.email 
		from customer c
		where upper(customer_name) like upper($1)
		order by c.customer_name
		limit $2 offset $3
	`
	f := "%" + filter + "%"
	rows, err := r.DB.QueryContext(ctx, s, f, PageSize, page*PageSize)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var customers []models.Customer
	for rows.Next() {
		var c models.Customer
		err := rows.Scan(
			&c.ID,
			&c.CustomerName,
			&c.Tel,
			&c.Email,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, c)
	}

	return &customers, nil
}

func (r *postgresCustomerRepo) SelectCustomerById(id int64) (*models.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	s := `
		select c.id, c.customer_name, c.tel, c.email, c.address_1, c.address_2, c.address_3, c.post_code 
		from customer c
		where c.id = $1
	`

	var c models.Customer
	row := r.DB.QueryRowContext(ctx, s, id)
	err := row.Scan(
		&c.ID,
		&c.CustomerName,
		&c.Tel,
		&c.Email,
		&c.Address1,
		&c.Address2,
		&c.Address3,
		&c.PostCode,
	)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *postgresCustomerRepo) CustomerUpdate(customer *models.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), DbTimeout)
	defer cancel()

	s := `
		update customer set 
			customer_name = $1, 
			tel = $2, 
			email = $3, 
			address_1 = $4, 
			address_2 = $5, 
			address_3 = $6, 
			post_code = $7,
			updated_at = $8
		where id = $9
	`

	_, err := r.DB.ExecContext(
		ctx,
		s,
		customer.CustomerName,
		customer.Tel,
		customer.Email,
		customer.Address1,
		customer.Address2,
		customer.Address3,
		customer.PostCode,
		customer.UpdatedAt,
		customer.ID,
	)

	return err
}
