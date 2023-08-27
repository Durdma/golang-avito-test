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

func (ss SlugsService) CreateNewSlug(slug *models.Slug) (*models.Slug, *models.ResponseError) {
	responseErr := ValidateSlug(slug)
	if responseErr != nil {
		return nil, responseErr
	}
	return ss.slugsRepository.CreateNewSlug(slug)
}

func (ss SlugsService) DelSlug(slugName string) (*models.Slug, *models.ResponseError) {
	responseErr := ValidateSlugName(slugName)
	if responseErr != nil {
		return nil, responseErr
	}
	return ss.slugsRepository.DelSlug(slugName)
}

func (ss SlugsService) PutSlug(slugName string, newSlug *models.Slug) (*models.Slug, *models.ResponseError) {
	responseErr := ValidateSlug(newSlug)
	if responseErr != nil {
		return nil, responseErr
	}

	responseErr = ValidateSlugName(slugName)
	if responseErr != nil {
		return nil, responseErr
	}

	return ss.slugsRepository.PutSlug(slugName, newSlug)
}
