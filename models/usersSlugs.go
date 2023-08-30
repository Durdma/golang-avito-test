package models

type UsersSlugs struct {
	UserUserID int    `gorm:"primaryKey;"`
	SlugSlugID int    `gorm:"primaryKey;"`
	CreatedAt  string `gorm:"not null" json:"created_at,omitempty"`
	UpdatedAt  string `gorm:"not null" json:"updated_at,omitempty"`
	DeletedAt  string `json:"deleted_at,omitempty"`
}
