package repositories

import (
	"avito-test/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type UsersRepository struct {
	dbHandler   *gorm.DB
	transaction *gorm.Tx
}

func NewUsersRepository(dbhandler *gorm.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbhandler,
	}
}

func (ur UsersRepository) AddUser(user *models.CreateUser) (*models.User, *models.ResponseError) {
	// err := ur.dbHandler.SetupJoinTable(&models.User{}, "ActiveSlugs", &models.UsersSlugs{})
	// if err != nil {
	// 	return nil, &models.ResponseError{
	// 		Messsage: err.Error(),
	// 		Status:   http.StatusInternalServerError,
	// 	}

	// }
	now := time.Now().Format(time.RFC3339)
	newUser := models.User{
		UserId:      user.UserId,
		CreatedAt:   now,
		UpdatedAt:   now,
		ActiveSlugs: []models.Slug{models.Slug{SlugId: 10101, SlugName: "TEST_ONE"}, models.Slug{SlugId: 202020, SlugName: "TEST_TWO"}},
	}

	result := ur.dbHandler.Create(&newUser)
	if result.Error != nil {
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	result = ur.dbHandler.Save(&newUser)
	if result.Error != nil {
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	return &models.User{
		UserId:      newUser.UserId,
		ActiveSlugs: newUser.ActiveSlugs,
		CreatedAt:   newUser.CreatedAt,
		UpdatedAt:   newUser.UpdatedAt,
	}, nil
}

func (ur UsersRepository) addSlugsToUser(userId string, listToAdd []*models.Slug) ([]*models.Slug, *models.ResponseError) {
	return nil, nil
}

func (ur UsersRepository) removeSlugsFromUser(userId string, listForRemove []*models.Slug) (*[]models.Slug, *models.ResponseError) {
	return nil, nil
}

func (ur UsersRepository) GetUser(userId string) (*models.User, *models.ResponseError) {
	var user models.User

	result := ur.dbHandler.First(&user, "user_id = ?", userId)
	if result.Error != nil {
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	return &user, nil
}
