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

	photoRepository := repositories.NewPhotoRepository(db)
	photoService := services.NewPhotoService(photoRepository)
	photoHandler := handlers.NewPhotoHandler(photoService)
	photoRouter := r.Group("/photos")

	{
		photoRouter.Use(middlewares.Authenthication())
		photoRouter.POST("/", photoHandler.Create)
		photoRouter.GET("/", photoHandler.GetAllPhotos)
		photoRouter.GET("/user", photoHandler.GetPhotosByUser)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), photoHandler.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), photoHandler.DeletePhoto)
	}

	commentRepository := repositories.NewCommentRepository(db)
	commentService := services.NewCommentService(commentRepository)
	commentHandler := handlers.NewCommentHandler(commentService)
	commentRouter := r.Group("/comments")

	{
		commentRouter.Use(middlewares.Authenthication())
		commentRouter.POST("/:photoId", commentHandler.Create)
		commentRouter.GET("/", commentHandler.GetComments)
	}

	r.Run()
}
