package routes

import (
	"backend/controllers"
	"backend/middleware"
	"github.com/gin-gonic/gin"
)

//
//func SetupRoutes(r *gin.Engine) {
//	userRoutes := r.Group("/users")
//	{
//		userRoutes.GET("/", controllers.GetUsers)
//	}
//}

func Init() error {
	router := gin.Default()
	adminRouter := router.Group("/admin")
	adminRouter.Use(middleware.AdminMiddleware())
	{
		adminRouter.POST("/createActivity", controllers.CreateActivity)
	}
	router.POST("/login", controllers.Login)
	router.POST("/register", controllers.Register)

	err := router.Run(":8080")
	if err != nil {
		return err
	}
	return nil
}
