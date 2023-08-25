package controllers

import (
	"avito-test/models"
	"avito-test/services"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsersController struct {
	usersService *services.UsersService
}

func NewUsersController(usersService *services.UsersService) *UsersController {
	return &UsersController{usersService: usersService}
}

func (uh UsersController) CreateUser(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create user request body!", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var user models.BaseUser
	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Println("Error while unmarshaling "+"creates user request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := uh.usersService.CreateNewUser(&user)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusCreated, response)
}

func (uh UsersController) GetUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	response, responseErr := uh.usersService.GetUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (uh UsersController) DelUser(ctx *gin.Context) {
	userId := ctx.Param("id")
	response, responseErr := uh.usersService.DelUser(userId)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}