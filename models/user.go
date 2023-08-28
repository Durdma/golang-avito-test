package models

import "time"

type User struct {
	Id int `json:"user_id" gorm:"primaryKey"`
	// SlugsListToAdd []*Slug `json:"slugs_list_to_add,omitempty"`
	// SlugsListToDel []*Slug `json:"slugs_list_to_del,omitempty"`
}

type CreateUser struct {
	Id             int       `json:"user_id" gorm:"primaryKey"`
	SlugsListToAdd []*string `json:"slugs_list_to_add,omitempty"`
	SlugsListToDel []*string `json:"slugs_list_to_del,omitempty"`
	CreatedAt      time.Time `json:"-"`
	UpdatedAt      time.Time `json:"-"`
	DeletedAt      time.Time `json:"-"`
}

type GetUser struct {
	Id          int       `json:"user_id" gorm:"primaryKey"`
	ActiveSlugs []*Slug   `json:"user_slugs_active,omitempty"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	DeletedAt   time.Time `json:"deleted_at,omitempty"`
}
