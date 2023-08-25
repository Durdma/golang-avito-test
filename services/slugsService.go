package services

import (
	"avito-test/models"
	"avito-test/repositories"
)

type SlugsService struct {
	usersRepository *repositories.UsersRepository
	slugsRepository *repositories.SlugsRepository
}

func NewSlugsService(usersRepository *repositories.UsersRepository, slugsRepository *repositories.SlugsRepository) *SlugsService {
	return &SlugsService{usersRepository: usersRepository, slugsRepository: slugsRepository}
}

func (ss SlugsService) CreateNewSlug(slug *models.BaseSlug) (*models.Slug, *models.ResponseError) {
	return nil, nil
}

func (ss SlugsService) DelSlug(slugId string) (*models.Slug, *models.ResponseError) {
	return nil, nil
}

func (ss SlugsService) PutSlug(slugId string, newSlug *models.BaseSlug) (*models.Slug, *models.ResponseError) {
	return nil, nil
}
