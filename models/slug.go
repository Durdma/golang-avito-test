package models

type BaseSlug struct {
	Id int `json:"slug_id"`
}

type Slug struct {
	BaseSlug
	Name string `json:"slug_name"`
}
