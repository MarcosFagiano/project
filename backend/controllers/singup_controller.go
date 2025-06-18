package controllers

import (
	"backend/dto"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Singup(c *gin.Context) {
	var singup dto.SingupDTO

	//Lamamos a BindJSON para unir la informacion del nuevo usuario
	//de JSON a user
	if err := c.ShouldBindJSON(&singup); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.Singup(singup)
	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}
