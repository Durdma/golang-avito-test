package services

import (
	"avito-test/models"
	"avito-test/repositories"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"
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

func (us UsersService) writeToCSV(history *[]models.History) (string, *models.ResponseError) {
	historyArr := make([][]string, len(*history))

	for i, val := range *history {
		historyArr[i] = append(historyArr[i], strconv.Itoa(val.UserUserId),
			val.SlugName, val.Operation, val.DateInfo[:])
		fmt.Println(historyArr[i])
	}

	filename := fmt.Sprintf("user_%v_history.csv", historyArr[0][0])

	file, e := os.Create(filename)
	defer file.Close()

	if e != nil {
		return "", &models.ResponseError{
			Messsage: "err csv",
			Status:   http.StatusInternalServerError,
		}
	}

	w := csv.NewWriter(file)

	e = w.Write([]string{"user_id", "slug_name", "operation", "date"})
	if e != nil {
		return "", &models.ResponseError{
			Messsage: "err write csv",
			Status:   http.StatusInternalServerError,
		}
	}

	e = w.WriteAll(historyArr)
	if e != nil {
		return "", &models.ResponseError{
			Messsage: "err write csv",
			Status:   http.StatusInternalServerError,
		}
	}

	return filename, nil
}

func (us UsersService) GetUserHistory(userId string) (string, *models.ResponseError) {
	history, err := us.usersRepository.GetUserHistory(userId)
	if err != nil {
		return "", err
	}

	filename, err := us.writeToCSV(history)

	return filename, err
}

// Подумать необходим ли данный поинт
// func (us UsersService) DelUser(userId string) (*models.User, *models.ResponseError) {
// 	return nil, nil
// }
