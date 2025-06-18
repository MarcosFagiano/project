package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	"backend/utils"
	"fmt"
)

func Login(repo *repository.UserRepository, loginDTO dto.LoginDTO) (*models.User, error) {
	userModel, err := repo.FindUserByEmail(loginDTO.Email)
	if err != nil {
		return nil, err // puede ser ErrRecordNotFound o DB error
	}

	// Validar el hash del password
	if utils.HashPasswordSHA256(loginDTO.Password) != userModel.Password {
		return nil, fmt.Errorf("invalid credentials")
	}

	return userModel, nil
}
