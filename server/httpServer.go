package server

import (
	"avito-test/controllers"
	"avito-test/repositories"
	"avito-test/services"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-co-op/gocron"
	"gorm.io/gorm"
)

type HttpServer struct {
	router          *gin.Engine
	usersController *controllers.UsersController
	slugsController *controllers.SlugsController
}

func runCron(taskOne func()) {
	s := gocron.NewScheduler(time.UTC)

	s.Every(1).Minutes().Do(taskOne)

	s.StartBlocking()
}

// For testing on early stages
func InitHttpServer(dbHandler *gorm.DB) HttpServer {

	slugsRepository := repositories.NewSlugsRepository(dbHandler)
	usersRepository := repositories.NewUsersRepository(dbHandler)

	usersService := services.NewUsersService(usersRepository, slugsRepository)
	slugsService := services.NewSlugsService(usersRepository, slugsRepository)

	usersController := controllers.NewUsersController(usersService)
	slugsController := controllers.NewSlugsController(slugsService)

	go func() {
		runCron(usersController.UpdateUserSlugsBySchedule)
	}()

	router := gin.Default()

	usersRouter := router.Group("/user")
	{
		usersRouter.POST("/", usersController.AddUser)
		usersRouter.GET("/:id", usersController.GetUser)
		usersRouter.GET("/:id/history", usersController.GetUserHistory)
	}

	slugsRouter := router.Group("/slug")
	{
		slugsRouter.POST("/", slugsController.CreateSlug)
		slugsRouter.DELETE("/:name", slugsController.DelSlug)
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
