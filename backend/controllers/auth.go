package controllers

import (
	"backend/models"
	"backend/services"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	authService *services.AuthService
}

func NewAuthController(authService *services.AuthService) *AuthController {
	return &AuthController{authService: authService}
}

func (ac *AuthController) Login(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userInput.Email == "" || userInput.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
		return
	}

	userLogin, err := ac.authService.Login(userInput.Email, userInput.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	if userLogin != nil {
		token, _ := utils.GenerateJWT(userLogin.Email, userLogin.Admin)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}
}

func (ac *AuthController) Register(c *gin.Context) {
	var userInput models.User
	if err := c.ShouldBind(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if userInput.Email == "" || userInput.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email and password required"})
		return
	}

	if err := ac.authService.Register(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}
