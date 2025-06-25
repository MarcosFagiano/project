package controllers

import (
	"backend/dto"
	"backend/mapper"
	"backend/models"
	"backend/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AdminController struct {
	adminService       *services.AdminService
	activityService    *services.ActivityService
	inscriptionService *services.InscriptionService
}

// Constructor
func NewAdminController(
	adminService *services.AdminService,
	activityService *services.ActivityService,
	inscriptionService *services.InscriptionService) *AdminController {
	return &AdminController{
		activityService:    activityService,
		inscriptionService: inscriptionService,
		adminService:       adminService,
	}
}

func (ac *AdminController) CreateActivity(c *gin.Context) {
	var activityDTO dto.Activity
	if err := c.ShouldBindJSON(&activityDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	activityModel := mapper.ActivityDTOToModel(activityDTO)
	if activityModel.StartDate.IsZero() || activityModel.EndDate.IsZero() {
		c.JSON(http.StatusBadRequest, gin.H{"error": "start_date and end_date are required and must be valid"})
		return
	}
	if err := ac.adminService.CreateActivity(activityModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"Success creation": activityModel})
}

func (ac *AdminController) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success creation": category})
}

func (ac *AdminController) DeleteActivity(c *gin.Context) {
	var activityDTO dto.Activity
	if err := c.ShouldBindJSON(&activityDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activityModel := mapper.ActivityDTOToModel(activityDTO)
	if err := ac.adminService.DeleteActivityByID(activityModel.ID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success deletion": activityModel})
}

func (ac *AdminController) UpdateActivity(c *gin.Context) {
	var activityDTO dto.Activity
	if err := c.ShouldBindJSON(&activityDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	activityModel := mapper.ActivityDTOToModel(activityDTO)
	if err := ac.adminService.UpdateActivity(activityModel); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"Success update": activityModel})
}
