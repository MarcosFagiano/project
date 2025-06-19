package repository

import (
	"backend/models"
	"errors"
	"gorm.io/gorm"
	"time"
)

type Repository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *Repository {
	return &Repository{DB: db}
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.DB.Where("email= ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
func (r *Repository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := r.DB.Preload("Inscriptions.Activities.Category").Find(&users).Error
	return users, err
}
func (r *Repository) CreateUser(user *models.User) error {

	return r.DB.Create(user).Error
}
func (r *Repository) CreateActivity(activity *models.Activity) error {
	return r.DB.Create(activity).Error
}

func (r *Repository) GetActivityByID(ID uint) (*models.Activity, error) {
	var activity models.Activity
	err := r.DB.Where("ID = ?", ID).First(&activity).Error
	return &activity, err
}

func (r *Repository) GetCategoryByID(ID uint) (*models.Category, error) {
	var category models.Category
	result := r.DB.Where("ID = ?", ID).First(&category)
	if result.Error != nil {
		return nil, result.Error
	}
	return &category, nil
}

func (r *Repository) CreateCategory(category *models.Category) error {
	return r.DB.Create(category).Error
}
func (r *Repository) ExistsActivity(description string, startDate, endDate time.Time, categoryID uint) (bool, error) {
	var activity models.Activity
	err := r.DB.Where("description = ? AND start_date = ? AND end_date = ? AND category_id = ?", description, startDate, endDate, categoryID).First(&activity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *Repository) ExistsCategoryByName(name string) (bool, error) {
	var category models.Category
	err := r.DB.Where("name = ?", name).First(&category).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}
