package models

type UsersSlugs struct {
	UserID int `gorm:"primaryKey;"`
	SlugID int `gorm:"primaryKey;"`
}
