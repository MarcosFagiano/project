package services

import (
	"backend/models"
	"backend/repository"
	"errors"
	"gorm.io/gorm"
)

type InscriptionService struct {
	inscriptionRepo *repository.InscriptionRepository
	activityRepo    *repository.ActivityRepository
	userRepo        *repository.UserRepository
}

// Constructor
func NewInscriptionService(
	inscriptionRepo *repository.InscriptionRepository,
	// activityRepo *repository.ActivityRepository,
	// userRepo *repository.UserRepository,
) *InscriptionService {
	return &InscriptionService{
		inscriptionRepo: inscriptionRepo,
		//activityRepo:    activityRepo,
		//userRepo:        userRepo,
	}
}

// Lógica de negocio para registrar inscripción
func (s *InscriptionService) RegisterToActivity(inscription models.Inscription) error {
	// Validaciones extra (cupo, fechas, existencia de usuario/actividad)
	quota, err := s.activityRepo.GetQuota(inscription.ActivityID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("activity not found")
		}
		return err

	}
	if quota == 0 {
		return errors.New("quota is zero")
	}
	_, err = s.userRepo.GetUserById(inscription.UserID)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return err

	}
	// Validacion de que el usuario no este inscripto a la actividad
	_, err = s.inscriptionRepo.GetInscriptionByUserId(inscription.UserID, inscription.ActivityID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return s.inscriptionRepo.CreateInscription(inscription)
		}
		return err
	}

	// Usuario ya inscripto
	return errors.New("user already registered to this activity")
}
