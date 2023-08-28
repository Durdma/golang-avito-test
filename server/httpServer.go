package server

import (
	"avito-test/controllers"
	"avito-test/repositories"
	"avito-test/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type HttpServer struct {
	router          *gin.Engine
	usersController *controllers.UsersController
	slugsController *controllers.SlugsController
}

// For testing on early stages
func InitHttpServer(dbHandler *gorm.DB) HttpServer {
	slugsRepository := repositories.NewSlugsRepository(dbHandler)
	usersRepository := repositories.NewUsersRepository(dbHandler)

	usersService := services.NewUsersService(usersRepository, nil)
	slugsService := services.NewSlugsService(nil, slugsRepository)

	usersController := controllers.NewUsersController(usersService)
	slugsController := controllers.NewSlugsController(slugsService)

	router := gin.Default()

	usersRouter := router.Group("/user")
	{
		usersRouter.POST("/", usersController.AddUser)
		usersRouter.GET("/:id", usersController.GetUser)
		usersRouter.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PLUG for /user/:id DELETE method", "id": ctx.Param("id")})
		})
	}

	slugsRouter := router.Group("/slug")
	{
		slugsRouter.POST("/", slugsController.CreateSlug)
		slugsRouter.DELETE("/:name", slugsController.DelSlug)
		slugsRouter.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PLUG for /slug/:id PUT method", "id": ctx.Param("id")})
		})
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	return HttpServer{router: router, usersController: usersController, slugsController: slugsController}
}

func (hs *HttpServer) RunServer() {
	err := hs.router.Run()
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
