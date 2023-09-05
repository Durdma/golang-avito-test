package models

type User struct {
	UserId      int    `gorm:"primaryKey;uniqueIndex;not null" json:"user_id"`
	ActiveSlugs []Slug `gorm:"many2many:users_slugs;foreignkey:user_id;association_foreignkey:slug_slug_id;association_jointable_foreignkey:slug_slug_id;jointable_foreignkey:user_user_id;"`
	CreatedAt   string `gorm:"type:varchar(150);not null" json:"created_at,omitempty"`
	UpdatedAt   string `gorm:"type:varchar(150);not null" json:"updated_at,omitempty"`
}

type CreateUser struct {
	UserId         int                 `json:"user_id" validate:"required,gte=1"`
	SlugsListToAdd []map[string]string `json:"slugs_list_to_add,omitempty" validate:"omitempty,slugs_list_to_add"`
	SlugsListToDel []string            `jsonL:"slugs_list_to_del,omitempty" validate:"omitempty,slugs_list_to_del"`
}

// JSON FORMAT

// {
//     "user_id": 1,
//     "slugs_list_to_add": [
//         {"slug_name": "TEST_SLUG_100", "days": ""}
//     ],
//	   "slugs_list_to_del": ["TEST_SLUG_001"]
// }
