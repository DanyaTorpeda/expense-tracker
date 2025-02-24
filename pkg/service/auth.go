package service

import (
	expensetracker "expense-tracker"
	"expense-tracker/pkg/repository"

	"golang.org/x/crypto/bcrypt"
)

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) Create(input expensetracker.User) (int, error) {
	hashedPassword, err := generatePasswordHash(input.Password)
	if err != nil {
		return 0, err
	}
	user := expensetracker.User{
		Name:     input.Name,
		Username: input.Username,
		Password: hashedPassword,
	}
	return s.repo.Create(user)
}

func (s *AuthService) GetUser(username, password string) (int, error) {
	user, err := s.repo.GetUser(username)
	if err != nil {
		return 0, err
	}

	if err := comparePassword(user.Password, password); err != nil {
		return 0, err
	}

	return user.ID, nil
}

func (s *AuthService) GenerateToken() {
}

func comparePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func generatePasswordHash(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}
