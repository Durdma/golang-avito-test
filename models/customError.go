package models

type ResponseError struct {
	Messsage string `json:"message"`
	Status   int    `json:"-"`
}
