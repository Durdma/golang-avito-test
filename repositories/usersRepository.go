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

func (ur UsersRepository) addSlugsToUser(user *models.User, slugs []string, activeSlugs []string) *models.ResponseError {
	for _, slugName := range slugs {
		for _, activeSlugName := range activeSlugs {
			if slugName == activeSlugName {
				return &models.ResponseError{
					Messsage: fmt.Sprintf("Error in user Slugs list to add"+
						"user already have this active slug %s", slugName),
					Status: http.StatusBadRequest,
				}
			}
		}
	}

	for _, slug := range slugs {
		var slugToAdd models.Slug

		result := ur.dbHandler.First(&slugToAdd, "slug_name = ?", slug)
		if result.Error != nil {
			if result.Error.Error() == "record not found" {
				return &models.ResponseError{
					Messsage: fmt.Sprintf("slug %s not found", slug),
					Status:   http.StatusNotFound,
				}
			}
		}

		user.ActiveSlugs = append(user.ActiveSlugs, slugToAdd)
	}

	return nil
}

func (ur UsersRepository) delUserSlug(userId int, slugId int) *models.ResponseError {
	var userSlug models.UsersSlugs

	result := ur.dbHandler.Where("user_user_id = ? AND slug_slug_id = ? AND disabled <> 1", userId, slugId).Delete(&userSlug)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}
	if result.RowsAffected == 0 {
		return &models.ResponseError{
			Messsage: "Something gone wrong!!!",
			Status:   http.StatusInternalServerError,
		}
	}

	result = ur.dbHandler.Where("user_user_id = ? AND slug_slug_id = ? AND disabled = 1", userId, slugId).Find(&userSlug)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	userSlug.UpdatedAt = time.Now().Format(time.RFC3339)
	userSlug.DeletedAt = time.Now().Format(time.RFC3339)

	result = ur.dbHandler.Save(&userSlug)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	return nil
}

func (ur UsersRepository) delSlugsFromUser(user *models.User, slugs []string, activeSlugs []string) *models.ResponseError {
	for _, slugName := range slugs {
		fl := false
		for _, activeSlugName := range activeSlugs {
			if slugName == activeSlugName {
				fl = true
				break
			}
		}
		if !fl {
			return &models.ResponseError{
				Messsage: fmt.Sprintf("Error in user Slugs list to del \n user doesnt have this active slug %s", slugName),
				Status:   http.StatusBadRequest,
			}
		}
	}

	if len(activeSlugs) > 0 {
		for _, slugToDel := range slugs {
			for i, slug := range user.ActiveSlugs {
				fmt.Printf("slugTodel: %s      slugActive: %s\n", slugToDel, slug.SlugName)
				if slugToDel == slug.SlugName {
					err := ur.delUserSlug(user.UserId, slug.SlugId)
					if err != nil {
						return err
					}

					user.ActiveSlugs = append(user.ActiveSlugs[:i], user.ActiveSlugs[i+1:]...)

					break
				}
			}
		}
	}

	return nil
}

// TODO исправить удаление на soft_delete user_slugs
// TODO вытекает проблема и с добавлением без soft_delete
// TODO Править время удаления через scoped
func (ur UsersRepository) AddUser(user *models.CreateUser) (*models.User, *models.ResponseError) {
	var newUser models.User
	var userSlug models.UsersSlugs

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

		err := ur.addSlugsToUser(&newUser, user.SlugsListToAdd, activeSlugsNames)
		if err != nil {
			return nil, err
		}

		err = ur.delSlugsFromUser(&newUser, user.SlugsListToDel, activeSlugsNames)
		if err != nil {
			return nil, err
		}

		newUser.UpdatedAt = time.Now().Format(time.RFC3339)

		result := ur.dbHandler.Save(&newUser)
		if result.Error != nil {
			return nil, &models.ResponseError{
				Messsage: result.Error.Error(),
				Status:   http.StatusInternalServerError,
			}
		}

	} else {
		if len(user.SlugsListToDel) != 0 {
			return nil, &models.ResponseError{
				Messsage: "new user cant have a list of slugs to del",
				Status:   http.StatusBadRequest,
			}
		}

		newUser.UserId = user.UserId
		newUser.CreatedAt = time.Now().Format(time.RFC3339)
		newUser.UpdatedAt = time.Now().Format(time.RFC3339)

		err := ur.addSlugsToUser(&newUser, user.SlugsListToAdd, make([]string, 0))
		if err != nil {
			return nil, err
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
	}
	// ТАКИМ ОБРАЗОМ РАБОТАТЬ СО СМЕЖНОЙ ТАБЛИЦЕЙ
	for _, slug := range newUser.ActiveSlugs {
		for _, addedSlug := range user.SlugsListToAdd {
			if addedSlug == slug.SlugName {
				result := ur.dbHandler.Model(&userSlug).Where("user_user_id = ? AND slug_slug_id = ? AND disabled <> 1",
					newUser.UserId, slug.SlugId).Updates(models.UsersSlugs{SlugSlugID: slug.SlugId, UserUserID: newUser.UserId,
					CreatedAt: time.Now().Format(time.RFC3339),
					UpdatedAt: time.Now().Format(time.RFC3339), DeletedAt: "", Disabled: 0})
				if result.Error != nil {
					return nil, &models.ResponseError{
						Messsage: result.Error.Error(),
						Status:   http.StatusInternalServerError,
					}
				}

				fmt.Printf("%+v\n", userSlug)
				result = ur.dbHandler.Save(&userSlug)
				if result.Error != nil {
					return nil, &models.ResponseError{
						Messsage: result.Error.Error(),
						Status:   http.StatusInternalServerError,
					}
				}
			}
		}
	}

	return &newUser, nil
}

// WORKING!!! ADD!!!
// func (ur UsersRepository) addSlugsToUser(userId string, listToAdd []*models.Slug) ([]*models.Slug, *models.ResponseError) {
// 	activeSlugsNames := make([]string, len(newUser.ActiveSlugs))
// 		for _, slug := range newUser.ActiveSlugs {
// 			activeSlugsNames = append(activeSlugsNames, slug.SlugName)
// 			fmt.Printf("# %s", slug.SlugName)
// 		}

// 		for _, slugName := range user.SlugsListToAdd {
// 			for _, activeSlugName := range activeSlugsNames {
// 				if slugName == activeSlugName {
// 					return nil, &models.ResponseError{
// 						Messsage: fmt.Sprintf("Error in user Slugs list to add"+
// 							"user already have this active slug %s", slugName),
// 						Status: http.StatusBadRequest,
// 					}
// 				}
// 			}

// 			for _, slug := range user.SlugsListToAdd {
// 				var slugToAdd models.Slug

// 				result := ur.dbHandler.First(&slugToAdd, "slug_name = ?", slug)
// 				if result.Error != nil {
// 					if result.Error.Error() == "record not found" {
// 						return nil, &models.ResponseError{
// 							Messsage: fmt.Sprintf("slug %s not found", slug),
// 							Status:   http.StatusNotFound,
// 						}
// 					}
// 				}

// 				newUser.ActiveSlugs = append(newUser.ActiveSlugs, slugToAdd)
// 			}

// 			result := ur.dbHandler.Save(&newUser)
// 			if result.Error != nil {
// 				return nil, &models.ResponseError{
// 					Messsage: result.Error.Error(),
// 					Status:   http.StatusInternalServerError,
// 				}
// 			}

// 			return &newUser, nil

// 	return nil, nil
// }

// func (ur UsersRepository) removeSlugsFromUser(userId string, listForRemove []*models.Slug) (*[]models.Slug, *models.ResponseError) {
// 	return nil, nil
// }

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
