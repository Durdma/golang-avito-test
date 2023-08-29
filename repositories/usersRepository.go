package repositories

import (
	"avito-test/models"
	"fmt"
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

func (ur UsersRepository) userExists(userId int) (*models.User, bool) {
	var activeUser models.User

	result := ur.dbHandler.Limit(1).Find(&activeUser, "user_id = ?", userId)
	if result.RowsAffected == 0 {
		return nil, false
	}

	activeUser.UserId = userId

	var res []*models.Slug

	err := ur.dbHandler.Model(&activeUser).Association("ActiveSlugs").Find(&res)
	if err != nil {
		return nil, false
	}

	for _, v := range res {
		// fmt.Printf("%+v\n", v)
		activeUser.ActiveSlugs = append(activeUser.ActiveSlugs, *v)
	}

	return &activeUser, true
}

// TODO Добавить функцию обработки списков сегментов и ограничения на таблицы в БД
func (ur UsersRepository) AddUser(user *models.CreateUser) (*models.User, *models.ResponseError) {

	now := time.Now().Format(time.RFC3339)
	var newUser models.User

	if result, ok := ur.userExists(user.UserId); ok {
		newUser = *result
		// fmt.Println("++++++")
		// fmt.Printf("%+v\n", newUser)
		// fmt.Println("++++++")
		activeSlugsNames := make([]string, len(newUser.ActiveSlugs))
		for _, slug := range newUser.ActiveSlugs {
			activeSlugsNames = append(activeSlugsNames, slug.SlugName)
			fmt.Printf("# %s", slug.SlugName)
		}

		for _, slugName := range user.SlugsListToAdd {
			for _, activeSlugName := range activeSlugsNames {
				if slugName == activeSlugName {
					return nil, &models.ResponseError{
						Messsage: fmt.Sprintf("Error in user Slugs list to add"+
							"user already have this active slug %s", slugName),
						Status: http.StatusBadRequest,
					}
				}
			}
			// err := ur.dbHandler.Model(&newUser).Association("ActiveSlugs").Append(&models.Slug{SlugId: 3, SlugName: slugName})
			// 	if err != nil {
			// 		return nil, &models.ResponseError{
			// 			Messsage: fmt.Sprintf("Error in function for adding slugs to activeSlugs" +
			// 				err.Error()),
			// 			Status: http.StatusInternalServerError,
			// 		}
			// 	}
		}

		for _, slugName := range user.SlugsListToDel {
			fl := false
			for _, activeSlugName := range activeSlugsNames {
				if slugName == activeSlugName {
					fl = true
					break
				}
			}
			if !fl {
				return nil, &models.ResponseError{
					Messsage: fmt.Sprintf("Error in user Slugs list to del \n user doesnt have this active slug %s", slugName),
					Status:   http.StatusBadRequest,
				}
			}
		}

		for _, slug := range user.SlugsListToAdd {
			var slugToAdd models.Slug

			result := ur.dbHandler.First(&slugToAdd, "slug_name = ?", slug)
			if result.Error != nil {
				if result.Error.Error() == "record not found" {
					return nil, &models.ResponseError{
						Messsage: fmt.Sprintf("slug %s not found", slug),
						Status:   http.StatusNotFound,
					}
				}
			}

			newUser.ActiveSlugs = append(newUser.ActiveSlugs, slugToAdd)
		}

		result := ur.dbHandler.Save(&newUser)
		if result.Error != nil {
			return nil, &models.ResponseError{
				Messsage: result.Error.Error(),
				Status:   http.StatusInternalServerError,
			}
		}

		return &newUser, nil

	} else {
		if len(user.SlugsListToDel) != 0 {
			return nil, &models.ResponseError{
				Messsage: "new user cant have a list of slugs to del",
				Status:   http.StatusBadRequest,
			}
		}
		newUser.UserId = user.UserId
		newUser.CreatedAt = now
		newUser.UpdatedAt = now

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
		//найти slugs в БД
		//Добавить юзера в БД
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
		fmt.Println("====")
		fmt.Println(result.Error)
		fmt.Println("====")
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	fmt.Printf("%+v\n", result)

	return &user, nil
}
