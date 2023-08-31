package repositories

import (
	"avito-test/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type UsersRepository struct {
	dbHandler *gorm.DB
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

// func (ur UsersRepository) addUserSlug(userSlug *models.UsersSlugs, userId int, slugId int) *models.ResponseError {
// 	result := ur.dbHandler.Model(&userSlug).Where("user_user_id = ? AND slug_slug_id = ?",
// 		userId, slugId).Updates(models.UsersSlugs{SlugSlugID: slugId, UserUserID: userId,
// 		CreatedAt: time.Now().Format(time.DateTime),
// 		UpdatedAt: time.Now().Format(time.DateTime), DeletedAt: ""})
// 	if result.Error != nil {
// 		return &models.ResponseError{
// 			Messsage: result.Error.Error(),
// 			Status:   http.StatusInternalServerError,
// 		}
// 	}

// 	result = ur.dbHandler.Save(&userSlug)
// 	if result.Error != nil {
// 		return &models.ResponseError{
// 			Messsage: result.Error.Error(),
// 			Status:   http.StatusInternalServerError,
// 		}
// 	}

// 	err := addUserOperation(ur.dbHandler, true, userSlug)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (ur UsersRepository) addUserSlug2(userSlug *models.UsersSlugs, userId int, slugId int, until string) *models.ResponseError {
	result := ur.dbHandler.Model(&userSlug).Where("user_user_id = ? AND slug_slug_id = ?",
		userId, slugId).Updates(models.UsersSlugs{SlugSlugID: slugId, UserUserID: userId,
		CreatedAt: time.Now().Format(time.DateTime),
		UpdatedAt: time.Now().Format(time.DateTime), DeletedAt: "", UntilDate: until})
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	result = ur.dbHandler.Save(&userSlug)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	err := addUserOperation(ur.dbHandler, true, userSlug)
	if err != nil {
		return err
	}

	return nil
}

func (ur UsersRepository) delUserSlug(userId int, slugId int) *models.ResponseError {
	var userSlug models.UsersSlugs

	result := ur.dbHandler.Model(&userSlug).Clauses(
		clause.Returning{Columns: []clause.Column{{Name: "slug_slug_id"},
			{Name: "user_user_id"}}}).Where("user_user_id = ? AND slug_slug_id = ?", userId, slugId).Delete(&userSlug)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	err := addUserOperation(ur.dbHandler, false, &userSlug)
	if err != nil {
		return err
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

func (ur UsersRepository) addSlugsToUser2(user *models.User, slugs []map[string]string, activeSlugs []string) *models.ResponseError {
	for _, slugName := range slugs {
		for _, activeSlugName := range activeSlugs {
			if slugName["slug_name"] == activeSlugName {
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

		result := ur.dbHandler.First(&slugToAdd, "slug_name = ?", slug["slug_name"])
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

func (ur UsersRepository) AddUserV2(user *models.CreateUser) (*models.User, *models.ResponseError) {
	var newUser models.User
	var userSlug models.UsersSlugs

	if result, ok := ur.userExists(user.UserId); ok {
		newUser = *result

		activeSlugsNames := make([]string, len(newUser.ActiveSlugs))
		for _, slug := range newUser.ActiveSlugs {
			activeSlugsNames = append(activeSlugsNames, slug.SlugName)
		}

		err := ur.addSlugsToUser2(&newUser, user.SlugsListToAdd, activeSlugsNames)
		if err != nil {
			return nil, err
		}

		err = ur.delSlugsFromUser(&newUser, user.SlugsListToDel, activeSlugsNames)
		if err != nil {
			return nil, err
		}

		newUser.UpdatedAt = time.Now().Format(time.DateTime)

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
		newUser.CreatedAt = time.Now().Format(time.DateTime)
		newUser.UpdatedAt = time.Now().Format(time.DateTime)

		err := ur.addSlugsToUser2(&newUser, user.SlugsListToAdd, make([]string, 0))
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
	}

	for _, slug := range newUser.ActiveSlugs {
		for _, addedSlug := range user.SlugsListToAdd {
			if addedSlug["slug_name"] == slug.SlugName {
				if addedSlug["days"] != "" {
					days, err := strconv.Atoi(addedSlug["days"]) // исправить на это конвертирование
					if err != nil {
						return nil, &models.ResponseError{
							Messsage: "incorrect id. id must be an uint",
							Status:   http.StatusBadRequest,
						}
					}
					days = 24 * days

					dateUntil := time.Now().Add(time.Hour * time.Duration(days)).Format(time.DateTime)
					result := ur.addUserSlug2(&userSlug, newUser.UserId, slug.SlugId, dateUntil) //REFACTOR userSLUG EMPTY
					if result != nil {
						return nil, result
					}
				} else {
					result := ur.addUserSlug2(&userSlug, newUser.UserId, slug.SlugId, "")
					if result != nil {
						return nil, result
					}
				}
			}
		}
	}

	return &newUser, nil
}

// func (ur UsersRepository) AddUser(user *models.CreateUser) (*models.User, *models.ResponseError) {
// 	var newUser models.User
// 	var userSlug models.UsersSlugs

// 	if result, ok := ur.userExists(user.UserId); ok {
// 		newUser = *result
// 		// fmt.Println("++++++")
// 		// fmt.Printf("%+v\n", newUser)
// 		// fmt.Println("++++++")
// 		activeSlugsNames := make([]string, len(newUser.ActiveSlugs))
// 		for _, slug := range newUser.ActiveSlugs {
// 			activeSlugsNames = append(activeSlugsNames, slug.SlugName)
// 			fmt.Printf("# %s", slug.SlugName)
// 		}

// 		err := ur.addSlugsToUser(&newUser, user.SlugsListToAdd, activeSlugsNames)
// 		if err != nil {
// 			return nil, err
// 		}

// 		err = ur.delSlugsFromUser(&newUser, user.SlugsListToDel, activeSlugsNames)
// 		if err != nil {
// 			return nil, err
// 		}

// 		newUser.UpdatedAt = time.Now().Format(time.DateTime)

// 		result := ur.dbHandler.Save(&newUser)
// 		if result.Error != nil {
// 			return nil, &models.ResponseError{
// 				Messsage: result.Error.Error(),
// 				Status:   http.StatusInternalServerError,
// 			}
// 		}

// 	} else {
// 		if len(user.SlugsListToDel) != 0 {
// 			return nil, &models.ResponseError{
// 				Messsage: "new user cant have a list of slugs to del",
// 				Status:   http.StatusBadRequest,
// 			}
// 		}

// 		newUser.UserId = user.UserId
// 		newUser.CreatedAt = time.Now().Format(time.DateTime)
// 		newUser.UpdatedAt = time.Now().Format(time.DateTime)

// 		err := ur.addSlugsToUser(&newUser, user.SlugsListToAdd, make([]string, 0))
// 		if err != nil {
// 			return nil, err
// 		}

// 		result := ur.dbHandler.Create(&newUser)
// 		if result.Error != nil {
// 			return nil, &models.ResponseError{
// 				Messsage: result.Error.Error(),
// 				Status:   http.StatusInternalServerError,
// 			}
// 		}

// 		// result = ur.dbHandler.Save(&newUser)
// 		// if result.Error != nil {
// 		// 	return nil, &models.ResponseError{
// 		// 		Messsage: result.Error.Error(),
// 		// 		Status:   http.StatusInternalServerError,
// 		// 	}
// 		// }
// 	}
// 	// ТАКИМ ОБРАЗОМ РАБОТАТЬ СО СМЕЖНОЙ ТАБЛИЦЕЙ
// 	for _, slug := range newUser.ActiveSlugs {
// 		for _, addedSlug := range user.SlugsListToAdd {
// 			if addedSlug == slug.SlugName {
// 				err := ur.addUserSlug(&userSlug, newUser.UserId, slug.SlugId)
// 				if err != nil {
// 					return nil, err
// 				}
// 			}
// 		}
// 	}

// 	return &newUser, nil
// }

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

// TODO ПОДПРАВИТЬ МЕТОД
func (ur UsersRepository) GetUser(userId string) (*models.User, *models.ResponseError) {
	var user *models.User

	userIdToGet, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		return nil, &models.ResponseError{
			Messsage: "incorrect id. id must be an uint",
			Status:   http.StatusBadRequest,
		}
	}

	user, exist := ur.userExists(int(userIdToGet))
	if exist {
		return user, nil
	}

	return nil, &models.ResponseError{
		Messsage: fmt.Sprintf("user with id %v not found", userId),
		Status:   http.StatusNotFound,
	}
}

// func (ur UsersRepository) addUserOperation(flag bool, userSlug *models.UsersSlugs) *models.ResponseError {
// 	var history models.History

// 	type slugName struct {
// 		SlugName string
// 	}

// 	var slug slugName

// 	history.UserUserId = userSlug.UserUserID

// 	if flag {
// 		history.Operation = "Added to Segment"
// 		history.DateInfo = userSlug.CreatedAt

// 		history.SlugName = slug.SlugName

// 	} else {
// 		history.Operation = "Deleted from Segment"
// 		history.DateInfo = time.Now().Format(time.RFC3339)
// 	}

// 	result := ur.dbHandler.Model(&models.Slug{}).Where("slug_id = ?", userSlug.SlugSlugID).First(&slug)
// 	if result.Error != nil {
// 		if result.Error.Error() == "record not found" {
// 			return &models.ResponseError{
// 				Messsage: result.Error.Error(),
// 				Status:   http.StatusNotFound,
// 			}
// 		}
// 	}

// 	history.SlugName = slug.SlugName

// 	result = ur.dbHandler.Save(&history)
// 	if result.Error != nil {
// 		return &models.ResponseError{
// 			Messsage: result.Error.Error(),
// 			Status:   http.StatusInternalServerError,
// 		}
// 	}

// 	return nil
// }

func (ur UsersRepository) GetUserHistory(userId string) (*[]models.History, *models.ResponseError) {
	var history []models.History

	userIdToGet, err := strconv.ParseInt(userId, 0, 0)
	if err != nil {
		return nil, &models.ResponseError{
			Messsage: "incorrect id. id must be an uint",
			Status:   http.StatusBadRequest,
		}
	}

	res := ur.dbHandler.Model(&models.History{}).Where("user_user_id = ?", userIdToGet).Find(&history)
	if res.Error != nil {
		return nil, &models.ResponseError{
			Messsage: "here history error",
			Status:   http.StatusInternalServerError,
		}
	}

	//fmt.Printf("here: \n %+v\n", history)
	return &history, nil
}

func (ur UsersRepository) UpdateUserSlugsBySchedule() *models.ResponseError {
	type userSlugTime struct {
		UserUserID int
		SlugSlugID int
		UntilDate  string
	}

	var records []userSlugTime

	fmt.Println("We are here")

	result := ur.dbHandler.Model(&models.UsersSlugs{}).Where("until_date <> ''").Find(&records)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: "Error on UpdateUserSlugsBySchedule" + result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	//fmt.Println(records)

	for _, rec := range records {
		untilDate, err := time.Parse(time.DateTime, rec.UntilDate)
		if err != nil {
			return &models.ResponseError{
				Messsage: err.Error(),
				Status:   http.StatusInternalServerError,
			}
		}

		// fmt.Printf("time now: %v\n", time.Now().UTC())
		// fmt.Printf("rec time: %v\n", untilDate)
		// fmt.Printf("eq: %v\n", time.Now().UTC().Add(3*time.Hour).Equal(untilDate))
		// fmt.Printf("gt: %v\n", time.Now().UTC().Add(3*time.Hour).After((untilDate)))
		// fmt.Printf("res: %v\n", time.Now().UTC().Add(3*time.Hour).Equal(untilDate) || time.Now().UTC().Add(3*time.Hour).After((untilDate)))
		if time.Now().UTC().Add(3*time.Hour).Equal(untilDate) || time.Now().UTC().Add(3*time.Hour).After((untilDate)) {
			err := ur.delUserSlug(rec.UserUserID, rec.SlugSlugID)
			if err != nil {
				return err
			}
			//fmt.Println("We are indel zone")
		}
	}

	return nil
}
