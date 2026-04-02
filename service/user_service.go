package service

import (
	"errors"
	"time"
	"todoapi/models"
	"todoapi/repository"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Login(email, password string) (models.LoginResponse, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo}
}

// Secret key for signing (you should definitely move this to your .env file)
const jwtSecret = "your_secret_key_change_this"

func (s *userServiceImpl) Login(email, password string) (models.LoginResponse, error) {
	// 1. Find user by email
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return models.LoginResponse{}, errors.New("invalid email or password")
	}

	// 2. Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return models.LoginResponse{}, errors.New("invalid email or password")
	}

	// 3. Update LastLogin
	now := time.Now()
	user.LastLogin = &now
	_ = s.repo.Update(&user) // Update last login in background

	// 4. Generate JWT Token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID.String(),
		"email":   user.Email,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return models.LoginResponse{}, err
	}

	return models.LoginResponse{
		Token: tokenString,
		User:  user,
	}, nil
}
