package service

import (
	"errors"
	expensetracker "expense-tracker"
	"expense-tracker/pkg/repository"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserID int `json:"user_id"`
}

const (
	tokenTTL = time.Hour * 12
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

func (s *AuthService) GenerateToken(id int) (string, error) {
	claims := tokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenTTL)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
		UserID: id,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s *AuthService) ParseToken(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	if err != nil {
		return 0, nil
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok || !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserID, nil
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
