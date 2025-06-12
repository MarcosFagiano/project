package routes

import (
	"github.com/gin-gonic/gin"
	"myapp/controllers"
)

func SetupRoutes(r *gin.Engine) {
	userRoutes := r.Group("/users")
	{
		userRoutes.GET("/", controllers.GetUsers)
	}
}
