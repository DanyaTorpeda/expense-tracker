package repository

import (
	expensetracker "expense-tracker"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	Create(input expensetracker.User) (int, error)
	GetUser(username string) (*expensetracker.User, error)
}

type Expense interface {
	Create(userId int, input expensetracker.Expense) (int, error)
	GetAll(userId int) ([]expensetracker.Expense, error)
}

type Repository struct {
	Authorization
	Expense
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Expense:       NewExpensePostgres(db),
	}
}
