package model

import "time"

type CreateUser struct {
	Name   string `form:"name" json:"name" binding:"required"`
	House  House  `form:"house" json:"house" binding:"required,isValidHouseNumber"`
	Mobile string `form:"mobile" json:"mobile" binding:"required" validate:"e164"`
	Email  string `form:"email" json:"email" binding:"required" validate:"email"`
}

type UserResponse struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Mobile string `json:"mobile"`
	House  House  `json:"house"`
}

type UpdateUser struct {
	Name   *string `json:"name" binding:"omitempty,omitnil"`
	House  *House  `json:"house" binding:"omitempty,omitnil,isValidHouseNumber"`
	Mobile *string `json:"mobile" binding:"omitempty,omitnil"`
	Email  *string `json:"email" binding:"omitempty,omitnil"`
}

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	House     House     `json:"house"`
	Mobile    string    `json:"mobile"`
	Email     string    `json:"email"`
	IsDeleted bool      `json:"isDeleted"`
	CreatedAt time.Time `json:"createdAt"`
}
