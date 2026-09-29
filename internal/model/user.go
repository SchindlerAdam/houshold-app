package model

type CreateUser struct {
	Name   string `form:"name" json:"name" binding:"required"`
	House  House  `form:"house" json:"house" binding:"required,isValidHouseNumber"`
	Mobile string `form:"mobile" json:"mobile" binding:"required" validate:"e164"`
	Email  string `form:"email" json:"email" binding:"required" validate:"email"`
}

type UpdatedUser struct {
	Name   *string `json:"name"`
	House  *House  `json:"house" binding:"isValidHouseNumber"`
	Mobile *string `json:"mobile"`
	Email  *string `json:"email"`
}

type User struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	House  House  `json:"house"`
	Mobile string `json:"mobile"`
	Email  string `json:"email"`
}
