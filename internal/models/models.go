package models

import "time"

type Customer struct {
	ID           int64
	CustomerName string
	Tel          string
	Email        string
	Address1     string
	Address2     string
	Address3     string
	PostCode     string
	Success      string
	Error        string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
