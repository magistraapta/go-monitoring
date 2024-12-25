package router

import (
	"go-backend/internal/controller"
	"go-backend/internal/metrics"
	"go-backend/internal/repository"
	service "go-backend/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	r.GET("/metrics", gin.WrapH(metrics.MetricsHandler))

	userRepo := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userController := controller.NewUserController(userService)

	userRoutes := r.Group("users")
	{
		userRoutes.POST("/", userController.CreateUser)
		userRoutes.GET("/", userController.GetAllUsers)
		userRoutes.DELETE("/:id", userController.DeleteUser)
		userRoutes.PUT("/:id", userController.UpdateUser)
		userRoutes.GET("/:id", userController.GetUserByID)
	}

	return r
}
