package services

import (
	"backend/repository"
)

type ActivityService struct {
	inscriptionRepo *repository.InscriptionRepository
	activityRepo    *repository.ActivityRepository
	userRepo        *repository.UserRepository
}

func NewActivityService(repo *repository.ActivityRepository) *ActivityService {
	return &ActivityService{activityRepo: repo}
}
