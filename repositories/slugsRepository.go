package repositories

import "avito-test/models"

type SlugsRepository struct {
}

func (sr SlugsRepository) CreateNewSlug(*models.Slug) (*models.Slug, *models.ResponseError) {
	return nil, nil
}

func (sr SlugsRepository) DelSlug(slugName string) (*models.Slug, *models.ResponseError) {
	return nil, nil
}

func (sr SlugsRepository) PutSlug(slugName string, newSlug *models.Slug) (*models.Slug, *models.ResponseError) {
	return nil, nil
}
