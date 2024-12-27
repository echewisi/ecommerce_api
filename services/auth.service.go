package services

import (
	"errors"
	"github.com/echewisi/ecommerce_api/models"
	"github.com/echewisi/ecommerce_api/repositories"
	"github.com/echewisi/ecommerce_api/utils"
)

type AuthService struct {
	UserRepo  *repositories.UserRepository
	JWTSecret string
}

// NewAuthService creates a new instance of AuthService
func NewAuthService(userRepo *repositories.UserRepository, jwtSecret string) *AuthService {
	return &AuthService{UserRepo: userRepo, JWTSecret: jwtSecret}
}

// RegisterUser registers a new user
func (s *AuthService) RegisterUser(email, password string, isAdmin bool) (*models.User, error) {
	// Check if the email already exists
	_, err := s.UserRepo.FindUserByEmail(email)
	if err == nil {
		return nil, errors.New("email already exists")
	}

	// Hash the password
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return nil, err
	}

	// Create the user
	user := &models.User{Email: email, Password: hashedPassword, IsAdmin: isAdmin}
	err = s.UserRepo.CreateUser(user)
	return user, err
}

// Login authenticates a user and generates a JWT token
func (s *AuthService) Login(email, password string) (string, error) {
	user, err := s.UserRepo.FindUserByEmail(email)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Verify the password
	err = utils.VerifyPassword(user.Password, password)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	// Generate a JWT token
	token, err := utils.GenerateToken(user.ID, user.IsAdmin, s.JWTSecret)
	return token, err
}
