package repositories

import (
	"avito-test/models"
	"database/sql"
)

type UsersRepository struct {
	dbHandler   *sql.DB
	transaction *sql.Tx
}

func NewUsersRepository(dbhandler *sql.DB) *UsersRepository {
	return &UsersRepository{
		dbHandler: dbhandler,
	}
}

func (ur UsersRepository) AddUser(user *models.User) (*models.User, *models.ResponseError) {
	return nil, nil
}

func (ur UsersRepository) addSlugsToUser(userId string, listToAdd []*models.Slug) ([]*models.Slug, *models.ResponseError) {
	return nil, nil
}

func (ur UsersRepository) removeSlugsFromUser(userId string, listForRemove []*models.Slug) (*[]models.Slug, *models.ResponseError) {
	return nil, nil
}

func (ur UsersRepository) GetUser(userId string) (*models.User, *models.ResponseError) {
	return nil, nil
}
