package controllers

import (
	"backend/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

var jwtKey = []byte("holis")

func Login(c *gin.Context) {
	var login dto.LoginDTO

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"email":   login.Email,
		"message": "success",
	})
}
