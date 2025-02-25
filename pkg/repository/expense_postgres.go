package repository

import (
	expensetracker "expense-tracker"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ExpensePostgres struct {
	db *sqlx.DB
}

func NewExpensePostgres(db *sqlx.DB) *ExpensePostgres {
	return &ExpensePostgres{db: db}
}

func (r *ExpensePostgres) Create(userId int, input expensetracker.Expense) (int, error) {
	query := fmt.Sprintf(`INSERT INTO %s (user_id, total, description, category, created_at) 
		values($1, $2, $3, $4, now()) RETURNING id`, expensesTable)
	row := r.db.QueryRow(query, userId, input.Total, input.Description, input.Category)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
