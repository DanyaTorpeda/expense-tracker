package repository

import (
	"errors"
	expensetracker "expense-tracker"
	"fmt"
	"strings"

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

func (r *ExpensePostgres) GetAll(userId int, sortBy string, order string) ([]expensetracker.Expense, error) {
	var expenses []expensetracker.Expense

	validSortFields := map[string]bool{"created_at": true, "total": true, "category": true}
	if !validSortFields[sortBy] {
		sortBy = "created_at"
	}

	if order != "ASC" && order != "DESC" {
		order = "DESC"
	}

	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 ORDER BY %s %s", expensesTable, sortBy, order)

	err := r.db.Select(&expenses, query, userId)
	if err != nil {
		return nil, err
	}

	return expenses, nil
}

func (r *ExpensePostgres) GetById(userId int, expenseId int) (*expensetracker.Expense, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 AND id = $2", expensesTable)
	var expense expensetracker.Expense
	err := r.db.Get(&expense, query, userId, expenseId)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (r *ExpensePostgres) Update(userId int, expenseId int, input expensetracker.Expense) error {
	//update table set column1=val, column2=val
	var argStr []string
	var argVal []interface{}
	argId := 1

	if input.Total != 0 {
		argStr = append(argStr, fmt.Sprintf("total = $%d", argId))
		argVal = append(argVal, input.Total)
		argId++
	}

	if input.Description != "" {
		argStr = append(argStr, fmt.Sprintf("description = $%d", argId))
		argVal = append(argVal, input.Description)
		argId++
	}

	if input.Category != "" {
		argStr = append(argStr, fmt.Sprintf("category = $%d", argId))
		argVal = append(argVal, input.Category)
		argId++
	}

	updateStr := strings.Join(argStr, ", ")
	query := fmt.Sprintf("UPDATE %s SET %s WHERE user_id = $%d AND id = $%d",
		expensesTable, updateStr, argId, argId+1)
	argVal = append(argVal, userId, expenseId)
	res, err := r.db.Exec(query, argVal...)
	if err != nil {
		return err
	}

	affectedRows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if affectedRows == 0 {
		return errors.New("expense not found, nothing updated")
	}

	return nil
}

func (r *ExpensePostgres) Delete(userId int, expenseId int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 AND id = $2", expensesTable)
	res, err := r.db.Exec(query, userId, expenseId)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("expense not found, nothing deleted")
	}

	return nil
}
