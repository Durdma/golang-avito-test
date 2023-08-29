package models

type User struct {
	UserId      int    `gorm:"primaryKey;uniqueIndex;not null" json:"user_id"`
	ActiveSlugs []Slug `gorm:"many2many:users_slugs;foreignkey:user_id;association_foreignkey:slug_slug_id;association_jointable_foreignkey:slug_slug_id;jointable_foreignkey:user_user_id;"`
	CreatedAt   string `gorm:"type:varchar(150);not null" json:"created_at,omitempty"`
	UpdatedAt   string `gorm:"type:varchar(150);not null" json:"updated_at,omitempty"`
}

type CreateUser struct {
	UserId         int      `json:"user_id"`
	SlugsListToAdd []string `json:"slugs_list_to_add,omitempty"`
	SlugsListToDel []string `json:"slugs_list_to_del,omitempty"`
	CreatedAt      string   `gorm:"type:varchar(150);not null" json:"created_at,omitempty"`
	UpdatedAt      string   `gorm:"type:varchar(150);not null" json:"updated_at,omitempty"`
}

type UsersSlugsInnerJoin struct {
	UserId int
	SlugId int
}
