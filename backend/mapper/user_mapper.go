package mapper

import (
	"backend/dto"
	"backend/models"
)

// DTO -> Model (para crear)
func UserDTOToModel(dto dto.UserDTO) models.User {
	return models.User{
		GivenName:   dto.GivenName,
		FamilyName:  dto.FamilyName,
		PhoneNumber: dto.PhoneNumber,
		Email:       dto.Email,
	}
}

// Model -> DTO (para responder)
func UserModelToDTO(user models.User) dto.UserDTO {
	return dto.UserDTO{
		GivenName:   user.GivenName,
		FamilyName:  user.FamilyName,
		PhoneNumber: user.PhoneNumber,
		Email:       user.Email,
		Admin:       user.Admin,
	}
}
