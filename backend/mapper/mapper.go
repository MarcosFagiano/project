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
