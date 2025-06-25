package routes

import (
	"backend/dependencies"
	"backend/middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Init(dep *dependencies.Dependencies) error {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // origen de tu frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))
	adminRouter := router.Group("/admin")
	adminRouter.Use(middleware.AdminAuth())
	{
		adminRouter.POST("/createActivity", dep.AdminController.CreateActivity)
		adminRouter.POST("/createCategory", dep.AdminController.CreateCategory)
		adminRouter.DELETE("/deleteActivity", dep.AdminController.DeleteActivity)
		adminRouter.PUT("/updateActivity", dep.AdminController.UpdateActivity)
	}
	userRouter := router.Group("/user")
	userRouter.Use(middleware.UserAuth())
	{
		userRouter.POST("/activities/inscriptions", dep.UserController.RegisterActivity)
		userRouter.GET("/getActivityList", dep.UserController.GetActivityList)
		userRouter.GET("/getCategoryList", dep.UserController.GetCategoryList)
		userRouter.GET("/getActivityById", dep.UserController.GetActivityById)

	}

	router.POST("/login", dep.AuthController.Login)
	router.POST("/register", dep.AuthController.Register)

	return router.Run(":8080")
}
