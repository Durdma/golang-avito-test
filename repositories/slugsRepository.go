package repositories

import (
	"avito-test/models"
	"fmt"
	"net/http"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type SlugsRepository struct {
	dbHandler   *gorm.DB
	transaction *gorm.Tx
}

func NewSlugsRepository(dbHandler *gorm.DB) *SlugsRepository {
	return &SlugsRepository{dbHandler: dbHandler}
}

func (sr SlugsRepository) slugExists(slugName string) *models.ResponseError {
	var slug models.Slug

	result := sr.dbHandler.First(&slug, "slug_name = ?", slugName)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return &models.ResponseError{
				Messsage: fmt.Sprintf("slug %s not found", slugName),
				Status:   http.StatusNotFound,
			}
		}
	}

	return nil
}

// TODO Переделать сообщение об ошибки на internal server error
// TODO добавление повторяющихся записей, добавить проверку, что запись не существует иначе ошибка????
func (sr SlugsRepository) CreateNewSlug(slug *models.CreateSlug) (*models.Slug, *models.ResponseError) {
	now := time.Now().Format(time.RFC3339)
	newSlug := models.Slug{
		SlugName:  slug.SlugName,
		CreatedAt: now,
		UpdatedAt: now,
	}

	result := sr.dbHandler.Create(&newSlug)
	if result.Error != nil {
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}
	return &models.Slug{
		SlugId:    newSlug.SlugId,
		SlugName:  slug.SlugName,
		CreatedAt: newSlug.CreatedAt,
		UpdatedAt: newSlug.UpdatedAt,
		DeletedAt: newSlug.DeletedAt,
		Disabled:  newSlug.Disabled,
	}, nil
}

func (sr SlugsRepository) delUserSlug(slugId int) *models.ResponseError {
	var userSlug models.UsersSlugs

	result := sr.dbHandler.Model(&userSlug).Where("slug_slug_id = ?", slugId).Delete(&userSlug)
	if result.Error != nil {
		return &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	return nil
}

func (sr SlugsRepository) DelSlug(slugName string) (*models.Slug, *models.ResponseError) {
	var delSlug models.Slug

	err := sr.slugExists(slugName)
	if err != nil {
		return nil, err
	}

	result := sr.dbHandler.Model(&delSlug).Clauses(clause.Returning{
		Columns: []clause.Column{{Name: "slug_id"}, {Name: "slug_name"}}}).Where("slug_name = ? AND disabled <> 1", slugName).Delete(&delSlug)
	if result.Error != nil {
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}

	err = sr.delUserSlug(delSlug.SlugId)
	if err != nil {
		return nil, err
	}

	// result = sr.dbHandler.Model(&delSlug).Where("slug_name = ? and disabled = 1", slugName).Updates(
	// 	models.Slug{}
	// )

	// result := sr.dbHandler.Where("slug_name = ? AND disabled <> 1", slugName).Delete(&delSlug)
	// if result.Error != nil {
	// 	return nil, &models.ResponseError{
	// 		Messsage: result.Error.Error(),
	// 		Status:   http.StatusInternalServerError,
	// 	}
	// }
	// if result.RowsAffected == 0 {
	// 	return nil, &models.ResponseError{
	// 		Messsage: fmt.Sprintf("Slug %s not found", slugName),
	// 		Status:   http.StatusNotFound,
	// 	}
	// }

	// result = sr.dbHandler.Where("slug_name = ? AND disabled = 1", slugName).Find(&delSlug)
	// if result.Error != nil {
	// 	return nil, &models.ResponseError{
	// 		Messsage: result.Error.Error(),
	// 		Status:   http.StatusInternalServerError,
	// 	}
	// }

	// delSlug.UpdatedAt = time.Now().Format(time.RFC3339)
	// delSlug.DeletedAt = time.Now().Format(time.RFC3339)

	// result = sr.dbHandler.Save(&delSlug)
	// if result.Error != nil {
	// 	return nil, &models.ResponseError{
	// 		Messsage: result.Error.Error(),
	// 		Status:   http.StatusInternalServerError,
	// 	}
	// }

	fmt.Printf("%+v\n", delSlug)

	return &delSlug, nil
}

func (sr SlugsRepository) PutSlug(slugName string, newSlug *models.Slug) (*models.Slug, *models.ResponseError) {
	return nil, nil
}
