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

func (us UsersService) AddUser(user *models.CreateUser) (*models.User, *models.ResponseError) {
	//responseErr := ValidateUser(user)
	// if responseErr != nil {
	// 	return nil, responseErr
	// }

	return us.usersRepository.AddUser(user)
}

func (us UsersService) GetUser(userId string) (*models.User, *models.ResponseError) {
	// responseErr := ValidateUserId(userId)
	// if responseErr != nil {
	// 	return nil, responseErr
	// }

	return us.usersRepository.GetUser(userId)
}

func (us UsersService) GetUserHistory(userId string) (*[]models.History, *models.ResponseError) {
	return us.usersRepository.GetUserHistory(userId)
}

// Подумать необходим ли данный поинт
// func (us UsersService) DelUser(userId string) (*models.User, *models.ResponseError) {
// 	return nil, nil
// }
