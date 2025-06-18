package repository

import (
	"backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email= ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *UserRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("Inscriptions.Activities.Category").Find(&users).Error
	return users, err
}
func (r *UserRepository) Create(user *models.User) error {

	return r.DB.Create(user).Error
}
