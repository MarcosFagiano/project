package services

import (
	"backend/models"
	"backend/repository"
	"errors"
	"gorm.io/gorm"
)

type UserService struct {
	inscriptionRepo *repository.InscriptionRepository
	activityRepo    *repository.ActivityRepository
	userRepo        *repository.UserRepository
}

// Constructor
func NewUserService(userRepo *repository.UserRepository,
	activityRepo *repository.ActivityRepository,
	inscriptionRepo *repository.InscriptionRepository) *UserService {
	return &UserService{
		userRepo:        userRepo,
		inscriptionRepo: inscriptionRepo,
		activityRepo:    activityRepo,
	}
}

// Lógica de negocio para registrar inscripción
func (s *UserService) RegisterToActivity(inscription models.Inscription) error {

	return s.inscriptionRepo.CreateInscription(inscription)
}

func (s *UserService) GetActivityList() ([]models.Activity, error) {
	activityList, err := s.activityRepo.GetAllActivity()
	if err != nil {
		return nil, err
	}
	return activityList, nil
}

func (s *UserService) GetCategoryList() ([]models.Category, error) {
	categoryList, err := s.activityRepo.GetCategoryList()
	if err != nil {
		return nil, err
	}
	return categoryList, nil
}

func (s *UserService) GetActivityById(id string) (*models.Activity, error) {
	activity, err := s.activityRepo.GetActivityById(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("activity not found")
		}
		return nil, err
	}
	return activity, nil
}
