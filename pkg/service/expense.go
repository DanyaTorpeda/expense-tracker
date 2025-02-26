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

func (s *ExpenseService) GetAll(userId int, sortBy string, order string) ([]expensetracker.Expense, error) {
	return s.repo.GetAll(userId, sortBy, order)
}

func (s *ExpenseService) GetById(userId int, expenseId int) (*expensetracker.Expense, error) {
	return s.repo.GetById(userId, expenseId)
}

func (s *ExpenseService) Update(userId int, expenseId int, input expensetracker.Expense) error {
	return s.repo.Update(userId, expenseId, input)
}

func (s *ExpenseService) Delete(userId int, expenseId int) error {
	return s.repo.Delete(userId, expenseId)
}
