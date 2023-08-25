package services

import (
	"avito-test/models"
	"avito-test/repositories"
)

type UsersService struct {
	usersRepository *repositories.UsersRepository
	slugsRepository *repositories.SlugsRepository
}

func NewUsersService(usersRepository *repositories.UsersRepository, slugsRepository *repositories.SlugsRepository) *UsersService {
	return &UsersService{usersRepository: usersRepository, slugsRepository: slugsRepository}
}

func (us UsersService) CreateNewUser(user *models.BaseUser) (*models.BaseUser, *models.ResponseError) {
	return nil, nil
}

func (us UsersService) GetUser(userId string) (*models.GetUser, *models.ResponseError) {
	return nil, nil
}

func (us UsersService) DelUser(userId string) (*models.GetUser, *models.ResponseError) {
	return nil, nil
}
