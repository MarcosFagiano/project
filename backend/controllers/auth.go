package controllers

import (
	"backend/config"
	"backend/dto"
	"backend/mapper"
	"backend/repository"
	"backend/utils"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(c *gin.Context) {
	repo := repository.NewUserRepository(config.DB)

	var input dto.UserRegister
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !utils.EmailValidation(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	_, err := repo.FindUserByEmail(input.Email)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email already registered"})
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	hashedPassword := utils.HashPassword(input.Password)
	newUser := mapper.UserRegisterDTOToModel(input, hashedPassword)
	if err := repo.Create(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context) {
	repo := repository.NewUserRepository(config.DB)

	var input dto.UserLogin
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if !utils.EmailValidation(input.Email) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email"})
		return
	}

	user, err := repo.FindUserByEmail(input.Email)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Email not found"})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Database error"})
		return
	}

	if !utils.CheckPassword(input.Password, user.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid credentials"})
		return
	}

	token, _ := utils.GenerateJWT(user.Email, user.Admin)
	user.Token = token
	repo.DB.Save(&user)

	c.JSON(http.StatusOK, gin.H{"token": user.Token})
}
