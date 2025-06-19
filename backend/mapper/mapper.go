package mapper

import (
	"backend/dto"
	"backend/models"
)

func UserRegisterDTOToModel(dto dto.UserRegister, hashedPassword string) models.User {
	return models.User{
		GivenName:   dto.GivenName,
		FamilyName:  dto.FamilyName,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
		Password:    hashedPassword,
		Admin:       false, // el admin no lo define el cliente
	}
}
func ActivityDTOToModel(dto dto.Activity) models.Activity {
	return models.Activity{
		ID:          dto.ID,
		Description: dto.Description,
		StartDate:   dto.StartDate,
		EndDate:     dto.EndDate,
		Quota:       dto.Quota,
		CategoryID:  dto.CategoryID,
	}
}

func ActivityModelToDTO(model models.Activity) dto.Activity {
	return dto.Activity{
		ID:          model.ID,
		Description: model.Description,
		StartDate:   model.StartDate,
		EndDate:     model.EndDate,
		Quota:       model.Quota,
		CategoryID:  model.CategoryID,
	}
}
func CategoryModelToDTO(model models.Category) dto.Category {
	return dto.Category{
		ID:          model.ID,
		Name:        model.Name,
		Description: model.Description,
	}
}

func CategoryDTOToModel(dto dto.Category) models.Category {
	return models.Category{
		ID:          dto.ID,
		Name:        dto.Name,
		Description: dto.Description,
	}
}
