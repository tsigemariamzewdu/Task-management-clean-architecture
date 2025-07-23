package router

import (
	"task_management/Delivery/controllers"
	"task_management/Repositories"
	usecases "task_management/usecases"
	"task_management/db"
	infrastruture "task_management/infrastructure"

	"github.com/gin-gonic/gin"
)

func SetUpRoutes(router *gin.Engine) error {
	// Initialize services
	passwordService := infrastruture.PasswordServiceImpl{}
	jwtService := infrastruture.JWTServiceImpl{}

	// Initialize repository
	userRepo:= repositories.NewUserRepository(db.UserCollection) 
	taskRepo := repositories.NewTaskRepository(db.TaskCollection)
	

	// Initialize user usecase
	userUsecase := usecases.UserUseCase{
		UserRepo:        userRepo,
		PasswordService: passwordService,
		JWTService:      jwtService,
	}
	taskUsecase:=usecases.TaskUseCase{
		TaskRepo: taskRepo,
	}


	// Initialize controller
	userController := controllers.NewUserController(&userUsecase)
	taskController :=controllers.NewTaskController(&taskUsecase)

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
	adminTaskRoutes:=router.Group("/tasks")
	adminTaskRoutes.Use(infrastruture.AuthWithRole("Admin"))
	{
		adminTaskRoutes.PUT("/:id", taskController.UpdateTaskByID)
		adminTaskRoutes.DELETE("/:id", taskController.DeleteTaskByID)
		adminTaskRoutes.POST("/", taskController.AddTask)

	}
	adminUserRoutes:=router.Group("/admin")
	adminUserRoutes.Use(infrastruture.AuthWithRole("Admin"))
	{
		adminUserRoutes.POST("/promote/",userController.PromoteUser)
	}
	
	return nil
}
