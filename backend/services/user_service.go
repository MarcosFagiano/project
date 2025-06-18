package services

import (
	"backend/dto"
	"backend/models"
	"backend/repository"
	_ "backend/repository"
	"backend/utils"
	"fmt"
)

var repo *repository.UserRepository

func Login(loginDTO dto.LoginDTO) (*models.User, error) {
	userModel, err := repo.FindUserByEmail(loginDTO.Email)
	if err != nil {
		return nil, err // puede ser ErrRecordNotFound o DB error
	}

	// Validar el hash del password
	if utils.HashPasswordSHA256(loginDTO.Password) != userModel.PasswordHash {
		return nil, fmt.Errorf("invalid credentials")
	}

	return userModel, nil
}

func Singup(singupDTO dto.SingupDTO) error {
	userModel, err := repo.FindUserByEmail(singupDTO.Email)
	if err != nil {
		return err
	}

	err = repo.Create(userModel)
	if err != nil {
		return err
	}
	return nil
}
