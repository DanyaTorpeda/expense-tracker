package service

import (
	"errors"
	expensetracker "expense-tracker"
	"expense-tracker/pkg/repository"
)

type ExpenseService struct {
	repo repository.Expense
}

func NewExpenseService(repo repository.Expense) *ExpenseService {
	return &ExpenseService{repo: repo}
}

func (s *ExpenseService) Create(userId int, input expensetracker.Expense) (int, error) {
	if !input.Category.Validate() {
		return 0, errors.New("wrong category data")
	}
	return s.repo.Create(userId, input)
}

func (s *ExpenseService) GetAll(userId int) ([]expensetracker.Expense, error) {
	return s.repo.GetAll(userId)
}
