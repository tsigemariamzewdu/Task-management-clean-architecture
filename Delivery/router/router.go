package router

import (
	"task_management/Delivery/controllers"
	infrastruture "task_management/Infrastructure"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(
	router *gin.Engine,
	userController *controllers.UserController,
	taskController *controllers.TaskController,
) error {
	// Public routes
	router.POST("/register", userController.Register)
	router.POST("/login", userController.Login)
	router.POST("/logout", userController.Logout)
	router.POST("/promote", userController.PromoteUser)

	userRoutes := router.Group("/tasks")
	userRoutes.Use(infrastruture.AuthWithRole("Admin","User"))
	{
		userRoutes.GET("/", taskController.GetTasks)
		userRoutes.GET("/:id", taskController.GetTaskByID)
	}
	
	adminTaskRoutes := router.Group("/tasks")
	adminTaskRoutes.Use(infrastruture.AuthWithRole("Admin"))
	{
		adminTaskRoutes.PUT("/:id", taskController.UpdateTaskByID)
		adminTaskRoutes.DELETE("/:id", taskController.DeleteTaskByID)
		adminTaskRoutes.POST("/", taskController.AddTask)
	}
	
	adminUserRoutes := router.Group("/admin")
	adminUserRoutes.Use(infrastruture.AuthWithRole("Admin"))
	{
		adminUserRoutes.POST("/promote/", userController.PromoteUser)
	}
	
	return nil
}