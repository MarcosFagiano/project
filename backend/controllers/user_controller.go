package controllers

import (
	"myapp/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	users := []models.User{
		{ID: 1, Name: "Juan", Email: "juan@example.com"},
	}
	c.JSON(http.StatusOK, users)
}
