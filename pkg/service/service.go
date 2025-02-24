package service

import (
	expensetracker "expense-tracker"
	"expense-tracker/pkg/repository"
)

type Authorization interface {
	Create(input expensetracker.User) (int, error)
	GetUser(username, password string) (int, error)
	GenerateToken()
}

type Service struct {
	Authorization
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
	}
}
