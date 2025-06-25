package services

import (
	"backend/models"
	"backend/repository"
	"backend/utils"
	"errors"
	"gorm.io/gorm"
	"time"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

// Register un usuario (puede ser admin o user según el campo Admin)
func (s *AuthService) Register(user models.User) error {
	// Verificar si el email ya existe
	existing, err := s.userRepo.GetUserByEmail(user.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if existing != nil {
		return errors.New("email already registered")
	}

	user.Password = utils.HashPassword(user.Password)
	user.RegisterDateStart = time.Now()
	user.RegisterDateEnd = user.RegisterDateStart.AddDate(0, 1, 0)

	return s.userRepo.CreateUser(&user)
}

// Login
func (s *AuthService) Login(email, password string) (*models.User, error) {
	if !utils.EmailValidation(email) {
		return nil, errors.New("invalid email")
	}
	user, err := s.userRepo.GetUserByEmail(email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !utils.CheckPassword(password, user.Password) {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
