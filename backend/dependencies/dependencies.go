package dependencies

import (
	"backend/controllers"
	"backend/repository"
	"backend/services"
	"gorm.io/gorm"
)

type Dependencies struct {
	UserController  *controllers.UserController
	AdminController *controllers.AdminController
	AuthController  *controllers.AuthController
}

func BuildDependencies(db *gorm.DB) (*Dependencies, error) {

	// Crear repos
	userRepo := repository.NewUserRepository(db)
	activityRepo := repository.NewActivityRepository(db)
	inscriptionRepo := repository.NewInscriptionRepository(db)

	// Crear services
	userService := services.NewUserService(userRepo, activityRepo, inscriptionRepo)
	adminService := services.NewAdminService(activityRepo, inscriptionRepo)
	activityService := services.NewActivityService(activityRepo)
	inscriptionService := services.NewInscriptionService(inscriptionRepo)
	authService := services.NewAuthService(userRepo) // o lo que corresponda

	// Crear controllers
	userController := controllers.NewUserController(userService, activityService, inscriptionService)
	adminController := controllers.NewAdminController(adminService, activityService, inscriptionService)
	authController := controllers.NewAuthController(authService)

	return &Dependencies{
		UserController:  userController,
		AdminController: adminController,
		AuthController:  authController,
	}, nil
}
