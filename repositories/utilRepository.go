package repositories

import (
	"avito-test/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

func addUserOperation(db *gorm.DB, flag bool, userSlug *models.UsersSlugs) *models.ResponseError {
	var history models.History

	type slugName struct {
		SlugName string
	}

	var slug slugName

	history.UserUserId = userSlug.UserUserID

	if flag {
		history.Operation = "Added to Segment"
		history.DateInfo = userSlug.CreatedAt

		history.SlugName = slug.SlugName

	} else {
		history.Operation = "Deleted from Segment"
		history.DateInfo = time.Now().Format(time.DateTime)
	}

	result := db.Model(&models.Slug{}).Unscoped().Where("slug_id = ?", userSlug.SlugSlugID).First(&slug)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return &models.ResponseError{
				Messsage: result.Error.Error(),
				Status:   http.StatusNotFound,
			}
		}
	}

	history.SlugName = slug.SlugName

	result = db.Create(&history)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	return nil
}
