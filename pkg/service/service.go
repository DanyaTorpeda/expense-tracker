package service

import (
	expensetracker "expense-tracker"
	"expense-tracker/pkg/repository"
)

type Authorization interface {
	Create(input expensetracker.User) (int, error)
	GetUser(username, password string) (int, error)
	GenerateToken(id int) (string, error)
	ParseToken(tokenString string) (int, error)
}

type Expense interface {
	Create(userId int, input expensetracker.Expense) (int, error)
	GetAll(userId int) ([]expensetracker.Expense, error)
}

type Service struct {
	Authorization
	Expense
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		Expense:       NewExpenseService(repository.Expense),
	}
}
