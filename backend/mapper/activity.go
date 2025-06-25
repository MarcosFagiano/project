package mapper

import (
	"backend/dto"
	"backend/models"
)

func ActivityDTOToModel(activity dto.Activity) *models.Activity {
	return &models.Activity{
		ID:          activity.ID,
		CategoryID:  activity.CategoryID,
		StartDate:   activity.StartDate,
		EndDate:     activity.EndDate,
		Quota:       activity.Quota,
		Description: activity.Description,
	}
}

func ActivityModelToDTO(activity *models.Activity) *dto.Activity {
	return &dto.Activity{
		ID:          activity.ID,
		CategoryID:  activity.CategoryID,
		StartDate:   activity.StartDate,
		EndDate:     activity.EndDate,
		Quota:       activity.Quota,
		Description: activity.Description,
	}
}

func ActivityModelToDtoList(activityList []models.Activity) []dto.Activity {
	var activity []dto.Activity
	for _, m := range activityList {
		activity = append(activity, *ActivityModelToDTO(&m))
	}
	return activity
}
