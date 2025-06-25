package services

import (
	"backend/models"
	"backend/repository"
	"errors"
	"gorm.io/gorm"
)

type AdminService struct {
	inscriptionRepo *repository.InscriptionRepository
	activityRepo    *repository.ActivityRepository
}

// Constructor
func NewAdminService(
	activityRepo *repository.ActivityRepository,
	inscriptionRepo *repository.InscriptionRepository) *AdminService {
	return &AdminService{
		inscriptionRepo: inscriptionRepo,
		activityRepo:    activityRepo,
	}
}

func (as *AdminService) CreateActivity(activity *models.Activity) error {
	_, err := as.activityRepo.GetActivityByDescription(activity.Description)
	if err != nil {
		return err
	}
	return as.activityRepo.CreateActivity(activity)
}

func (as *AdminService) GetActivityByDescription(description string) (*models.Activity, error) {
	return as.activityRepo.GetActivityByDescription(description)
}

func (as *AdminService) GetAllActivity() ([]models.Activity, error) {
	return as.activityRepo.GetAllActivity()
}

func (as *AdminService) DeleteActivityByID(ID uint) error {
	if err := as.activityRepo.DeleteActivity(ID); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("activity not found")
		}
		return err
	}
	return nil
}

func (as *AdminService) UpdateActivity(activity *models.Activity) error {
	if err := as.activityRepo.UpdateActivity(activity); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("activity not found")
		}
		return err
	}
	return nil
}
