package models

type User struct {
	Id             int     `json:"user_id"`
	SlugsListToAdd []*Slug `json:"slugs_list_to_add,omitempty"`
	SlugsListToDel []*Slug `json:"slugs_list_to_del,omitempty"`
}
