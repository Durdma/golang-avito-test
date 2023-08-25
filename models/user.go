package models

type BaseUser struct {
	Id int `json:"user_id"`
}

type GetUser struct {
	BaseUser
	SlugsList []*Slug `json:"slugs_list"`
}
