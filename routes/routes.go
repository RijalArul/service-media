package routes

import (
	"service-media/databases"
	"service-media/handlers"
	"service-media/middlewares"
	"service-media/repositories"
	"service-media/services"

	"github.com/gin-gonic/gin"
)

func Routes() {
	r := gin.Default()
	db := databases.GetDB()
	userRepository := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepository)
	userHandler := handlers.NewUserHandler(userService)
	userRouter := r.Group("/users")

	{
		userRouter.POST("/register", userHandler.Register)
		userRouter.POST("/login", userHandler.Login)
		userRouter.Use(middlewares.Authenthication())
		userRouter.GET("/", userHandler.GetUser)
		userRouter.PUT("/", userHandler.UpdateUser)
		userRouter.DELETE("/", userHandler.DeleteUser)
	}

	r.Run()
}
