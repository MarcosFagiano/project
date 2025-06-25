package repository

import (
	"backend/models"
	"gorm.io/gorm"
)

type InscriptionRepository struct {
	db *gorm.DB
}

func NewInscriptionRepository(db *gorm.DB) *InscriptionRepository {
	return &InscriptionRepository{db: db}
}

func (r *InscriptionRepository) CreateInscription(inscription models.Inscription) error {
	return r.db.Create(&inscription).Error
}
func (r *InscriptionRepository) FindInscriptionByUserID(userId string) (*models.Inscription, error) {
	var inscription models.Inscription
	return &inscription, r.db.Where("id = ?", userId).First(&inscription).Error
}

func (r *InscriptionRepository) GetInscriptionByUserId(userId uint, activityId uint) (*models.Inscription, error) {
	var inscription models.Inscription
	return &inscription, r.db.Where("user_id = ? AND activity_id = ?", userId, activityId).First(&inscription).Error
}
