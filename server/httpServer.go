package server

import (
	"avito-test/controllers"
	"avito-test/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	router          *gin.Engine
	usersController *controllers.UsersController
}

// For testing on early stages
func InitHttpServer() HttpServer {
	usersService := services.NewUsersService(nil, nil)

	usersController := controllers.NewUsersController(usersService)

	router := gin.Default()

	usersRouter := router.Group("/user")
	{
		usersRouter.POST("/", usersController.CreateUser)
		usersRouter.GET("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PLUG for /user/:id GET method", "id": ctx.Param("id")})
		})
		usersRouter.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PLUG for /user/:id DELETE method", "id": ctx.Param("id")})
		})
	}

	slugsRouter := router.Group("/slug")
	{
		slugsRouter.POST("/", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "PLUG for /slug/ POST method")
		})
		slugsRouter.DELETE("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PLUG for /slug/:id DELETE method", "id": ctx.Param("id")})
		})
		slugsRouter.PUT("/:id", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "PLUG for /slug/:id PUT method", "id": ctx.Param("id")})
		})
	}

	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	return HttpServer{router: router, usersController: usersController}
}

func (hs *HttpServer) RunServer() {
	err := hs.router.Run()
	if err != nil {
		log.Fatalf("Error while starting HTTP server: %v", err)
	}
}
