package controllers

import (
	"avito-test/models"
	"avito-test/services"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SlugsController struct {
	slugsService *services.SlugsService
}

func NewSlugsController(slugsService *services.SlugsService) *SlugsController {
	return &SlugsController{slugsService: slugsService}
}

func (us SlugsController) CreateSlug(ctx *gin.Context) {
	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading create slug request body!", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var slug models.CreateSlug
	err = json.Unmarshal(body, &slug)
	if err != nil {
		log.Println("Error while unmarshaling "+"creates slug request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := us.slugsService.CreateNewSlug(&slug)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}
	fmt.Printf("response \n %+v\n", response)
	ctx.JSON(http.StatusCreated, response)
}

func (us SlugsController) DelSlug(ctx *gin.Context) {
	slugName := ctx.Param("name")
	response, responseErr := us.slugsService.DelSlug(slugName)
	if responseErr != nil {
		ctx.AbortWithStatusJSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}

func (us SlugsController) PutSlug(ctx *gin.Context) {
	slugId := ctx.Param("name")

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		log.Println("Error while reading put slug request body!", err)
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, err)
		return
	}

	var slug *models.Slug
	err = json.Unmarshal(body, &slug)
	if err != nil {
		log.Println("Error while unmarshaling "+"creates slug request body", err)
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	response, responseErr := us.slugsService.PutSlug(slugId, slug)
	if responseErr != nil {
		ctx.JSON(responseErr.Status, responseErr)
		return
	}

	ctx.JSON(http.StatusOK, response)
}
