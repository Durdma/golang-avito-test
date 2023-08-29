package models

import "gorm.io/plugin/soft_delete"

type UsersSlugs struct {
	UserUserID int                   `gorm:"primaryKey;"`
	SlugSlugID int                   `gorm:"primaryKey;"`
	CreatedAt  string                `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt  string                `gorm:"not null" json:"updated_at,omitempty"`
	DeletedAt  string                `json:"deleted_at,omitempty"`
	Disabled   soft_delete.DeletedAt `gorm:"softDelete:flag" json:"disabled,omitempty"`
}
