package mapper

import (
	"backend/dto"
	"backend/models"
)

func InscriptionDtoToModel(create dto.InscriptionCreate) models.Inscription {
	return models.Inscription{
		UserID:     create.UserID,
		ActivityID: create.ActivityID,
	}
}
