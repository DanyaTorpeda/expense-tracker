package repository

import (
	expensetracker "expense-tracker"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(input expensetracker.User) (int, error)
	GetUser(username string) (*expensetracker.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
