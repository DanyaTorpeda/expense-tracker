package repository

import (
	expensetracker "expense-tracker"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) Create(input expensetracker.User) (int, error) {
	query := fmt.Sprintf("INSERT INTO %s (name, username, password_hash) values($1, $2, $3) RETURNING id", usersTable)
	row := r.db.QueryRow(query, input.Name, input.Username, input.Password)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(username string) (*expensetracker.User, error) {
	query := fmt.Sprintf("SELECT * FROM %s WHERE username = $1", usersTable)
	var user expensetracker.User
	err := r.db.Get(&user, query, username)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
