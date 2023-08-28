package models

import (
	"time"

	"gorm.io/plugin/soft_delete"
)

// Посмотреть как будет работать
type Slug struct {
	SlugId    int                   `gorm:"primaryKey;not null;type:serial" json:"slug_id,omitempty"`
	SlugName  string                `gorm:"uniqueIndex;not null" json:"slug_name,omitempty"`
	CreatedAt time.Time             `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt time.Time             `gorm:"not null" json:"updated_at,omitempty"`
	DeletedAt time.Time             `json:"deleted_at,omitempty"`
	Disabled  soft_delete.DeletedAt `gorm:"softDelete:flag,DeletedAtField:DeletedAt" json:"disabled,omitempty"`
	Users     []User                `gorm:"many2many:users_slugs"`
}

type CreateSlug struct {
	SlugName  string    `json:"slug_name"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}

type UpdateSlug struct {
	SlugName  string    `json:"slug_name,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	DeletedAt time.Time `json:"deleted_at,omitempty"`
}
