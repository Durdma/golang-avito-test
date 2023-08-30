package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// Посмотреть как будет работать
type Slug struct {
	SlugId    int                   `gorm:"primaryKey;not null;type:serial" json:"slug_id,omitempty"`
	SlugName  string                `gorm:"uniqueIndex;not null" json:"slug_name,omitempty"`
	CreatedAt string                `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt string                `gorm:"not null" json:"updated_at,omitempty"`
	DeletedAt string                `json:"deleted_at,omitempty"`
	Disabled  soft_delete.DeletedAt `gorm:"softDelete:flag" json:"disabled,omitempty"`
	Users     []User                `json:"-" gorm:"many2many:users_slugs;foreignkey:slug_id;association_foreignkey:user_user_id;association_jointable_foreignkey:user_user_id;jointable_foreignkey:slug_slug_id;"`
}

type CreateSlug struct {
	SlugName  string    `json:"slug_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
