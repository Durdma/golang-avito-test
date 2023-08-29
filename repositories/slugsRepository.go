package repositories

import (
	"avito-test/models"
	"net/http"
	"time"

	"gorm.io/gorm"
)

type SlugsRepository struct {
	dbHandler   *gorm.DB
	transaction *gorm.Tx
}

func (sr SlugsRepository) SlugExists(slugName string) (*models.Slug, bool) {
	var slug models.Slug

	result := sr.dbHandler.First(&slug, "slug_name = ?", slugName)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, false
		}
	}

	return &models.Slug{
		SlugId:    slug.SlugId,
		SlugName:  slugName,
		CreatedAt: slug.CreatedAt,
		UpdatedAt: slug.UpdatedAt,
		DeletedAt: slug.DeletedAt,
		Disabled:  slug.Disabled,
	}, true
}

func NewSlugsRepository(dbHandler *gorm.DB) *SlugsRepository {
	return &SlugsRepository{dbHandler: dbHandler}
}

// TODO Добавить функцию обработки списков сегментов и ограничения на таблицы в БД
// TODO Переделать сообщение об ошибки на internal server error
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

func (sr SlugsRepository) DelSlug(slugName string) (*models.Slug, *models.ResponseError) {
	var response *models.Slug

	result := sr.dbHandler.Where("slug_name = ?", slugName).Delete(&response)
	if result.Error != nil {
		return nil, &models.ResponseError{
			Messsage: result.Error.Error(),
			Status:   http.StatusInternalServerError,
		}
	}
	return &models.Slug{
		SlugId:    response.SlugId,
		SlugName:  response.SlugName,
		CreatedAt: response.CreatedAt,
		UpdatedAt: response.UpdatedAt,
		DeletedAt: response.DeletedAt,
		Disabled:  response.Disabled,
	}, nil
}

func (sr SlugsRepository) PutSlug(slugName string, newSlug *models.Slug) (*models.Slug, *models.ResponseError) {
	return nil, nil
}
