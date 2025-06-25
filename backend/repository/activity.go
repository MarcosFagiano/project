package repository

import (
	"backend/models"
	"errors"
	"gorm.io/gorm"
)

type ActivityRepository struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) *ActivityRepository {
	return &ActivityRepository{db: db}
}

//metodos para ver, crear, eliminar y actualizar una actividad

func (r *ActivityRepository) CreateActivity(activity *models.Activity) error {
	return r.db.Create(activity).Error
}
func (r *ActivityRepository) GetActivityByDescription(description string) (*models.Activity, error) {
	var activity models.Activity
	err := r.db.Where("description = ?", description).First(&activity).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &activity, nil
}

func (r *ActivityRepository) GetAllActivity() ([]models.Activity, error) {
	var activities []models.Activity
	if err := r.db.Model(&models.Activity{}).Find(&activities).Error; err != nil {
		return nil, err
	}
	if len(activities) == 0 {
		return nil, nil
	}
	return activities, nil
}

// al estar relacionada la categoria con la actividad, voy a poner los metodos dentro de este repositorio
func (r *ActivityRepository) GetCategoryList() ([]models.Category, error) {
	var categories []models.Category
	if err := r.db.Model(&models.Category{}).Find(&categories).Error; err != nil {
		return nil, err
	}
	if len(categories) == 0 {
		return nil, nil
	}
	return categories, nil
}

// metodos para ver y actualizar los cupos el la actividad dependiendo de las inscripciones
func (r *ActivityRepository) GetQuota(id uint) (uint, error) {
	var activity models.Activity
	err := r.db.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return 0, err
	}
	return activity.Quota, nil
}

func (r *ActivityRepository) UpdateQuota(id uint, quota uint) error {
	var activity models.Activity
	err := r.db.Where("id = ?", id).First(&activity).Error
	if err != nil {
		return err
	}
	activity.Quota = quota
	return r.db.Save(&activity).Error
}

func (r *ActivityRepository) DeleteActivity(id uint) error {
	//se eeliminan las inscripciones asociadas a la actividad
	err := r.db.Where("activity_id = ?", id).Delete(&models.Inscription{}).Error
	if err != nil {
		return err
	}

	//se eliminan la actividad por id ( el front recibe la lista de actividades y se selecciona por id por simplicidad)
	err = r.db.Where("id = ?", id).Delete(&models.Activity{}).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *ActivityRepository) UpdateActivity(activity *models.Activity) error {
	return r.db.Model(&models.Activity{}).Where("id = ?", activity.ID).Updates(&activity).Error
}

func (r *ActivityRepository) GetActivityById(id string) (*models.Activity, error) {
	var activity models.Activity
	return &activity, r.db.Where("id = ?", id).First(&activity).Error
}
