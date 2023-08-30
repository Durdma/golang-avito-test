package models

type History struct {
	OperationId int    `gorm:"primaryKey;not null;type:serial" json:"-"`
	UserUserId  int    `gorm:"not null" json:"user_id"`
	SlugName    string `gorm:"type:varchar(150);not null" json:"slug_name"`
	Operation   string `gorm:"type:varchar(50);not null" json:"operation_name"`
	DateInfo    string `gorm:"type:varchar(150);not null" json:"date_info"`
}
