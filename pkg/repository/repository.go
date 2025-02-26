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
	GetAll(userId int, sortBy string, order string) ([]expensetracker.Expense, error)
	GetById(userId int, expenseId int) (*expensetracker.Expense, error)
	Update(userId int, expenseId int, input expensetracker.Expense) error
	Delete(userId int, expenseId int) error
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
