package model

type CreateUser struct {
	Name   string `json:"name"`
	House  House  `json:"house"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	House  House  `json:"house"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
