package controllers

import (
	"backend/dto"
	"backend/mapper"
	"backend/models"
	"backend/services"
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type UserController struct {
	userService        *services.UserService
	activityService    *services.ActivityService
	inscriptionService *services.InscriptionService
}

// Constructor
func NewUserController(
	userService *services.UserService,
	activityService *services.ActivityService,
	inscriptionService *services.InscriptionService) *UserController {
	return &UserController{
		activityService:    activityService,
		inscriptionService: inscriptionService,
		userService:        userService}
}

// Handler
func (uc *UserController) RegisterActivity(c *gin.Context) {
	var input dto.InscriptionCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	inscriptionModel := mapper.InscriptionDtoToModel(input)
	if err := uc.userService.RegisterToActivity(inscriptionModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Inscription created successfully"})
}

func (uc *UserController) GetActivityList(c *gin.Context) {
	activityListModel, err := uc.userService.GetActivityList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activityListDto := mapper.ActivityModelToDtoList(activityListModel)
	c.JSON(http.StatusOK, gin.H{"data": activityListDto})
}

func (uc *UserController) GetCategoryList(c *gin.Context) {
	categoryListModel, err := uc.userService.GetCategoryList()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	categoryListDto := mapper.CategoryModelToDtoList(categoryListModel)
	c.JSON(http.StatusOK, gin.H{"data": categoryListDto})
}

func (uc *UserController) GetActivityById(c *gin.Context) {
	var activityInput *models.Activity
	if err := c.ShouldBindJSON(&activityInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
	}

	activity, err := uc.userService.GetActivityById(strconv.Itoa(int(activityInput.ID)))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Activity not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": activity})
}
