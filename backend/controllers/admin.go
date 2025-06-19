package controllers

import (
	"backend/config"
	"backend/dto"
	"backend/mapper"
	"backend/repository"
	"github.com/gin-gonic/gin"
	"net/http"
)

func CreateActivity(c *gin.Context) {
	repo := repository.NewUserRepository(config.DB)

	var input dto.Activity
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := repo.ExistsActivity(input.Description, input.StartDate, input.EndDate, input.CategoryID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for duplicates"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Activity already exists"})
		return
	}

	newActivity := mapper.ActivityDTOToModel(input)
	if err = repo.CreateActivity(&newActivity); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Activity created": mapper.ActivityModelToDTO(newActivity)})
}

func CreateCategory(c *gin.Context) {
	repo := repository.NewUserRepository(config.DB)

	var input dto.Category
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	exists, err := repo.ExistsCategoryByName(input.Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking for duplicates"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Category already exists"})
		return
	}

	newCategory := mapper.CategoryDTOToModel(input)
	if err = repo.CreateCategory(&newCategory); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Category created": mapper.CategoryModelToDTO(newCategory)})
}
