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
	Register(email, password string) (models.LoginResponse, error)
}

type userServiceImpl struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userServiceImpl{repo}
}

const jwtSecret = "MyJwtSecretKey"

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
func (s *userServiceImpl) Register(email, password string) (models.LoginResponse, error) {
	// 1. Check if user already exists
	_, err := s.repo.FindByEmail(email)
	if err == nil {
		return models.LoginResponse{}, errors.New("user already exists")
	}

	// 2. Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return models.LoginResponse{}, err
	}

	// 3. Create the user object
	user := models.User{
		Email:        email,
		PasswordHash: string(hashedPassword),
	}

	// 4. Save to database
	err = s.repo.Create(&user)
	if err != nil {
		return models.LoginResponse{}, err
	}

	// 5. Automatically log in after registration
	return s.Login(email, password)
}
