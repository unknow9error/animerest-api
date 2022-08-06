package main

import (
	controllers "animerest/Controllers"
	database "animerest/Database"
	docs "animerest/docs"

	services "animerest/Services"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	dbContext := database.Connect()
	r := gin.Default()

	userService := &services.UserService{DbContext: dbContext}
	userController := &controllers.UserController{UserService: userService}

	animeService := &services.AnimeService{DbContext: dbContext}
	animeController := &controllers.AnimeController{AnimeService: animeService}

	docs.SwaggerInfo.BasePath = "/api/"
	api := r.Group("/api/")
	{
		user := api.Group("/user")
		{
			user.POST("/login", userController.Login)
			user.POST("/register", userController.Register)
		}
		anime := api.Group("/anime")
		{
			anime.GET("/", animeController.GetAll)
			anime.GET("/title", animeController.GetById)
		}
	}

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run(":8080")
}
