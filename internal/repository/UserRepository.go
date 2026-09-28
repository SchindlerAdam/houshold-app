package repository

import (
	"database/sql"
	"houshold-app/internal/model"
	"log"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (repository UserRepository) Create(user *model.CreateUser) (*model.User, error) {
	var createdUser model.User
	err := repository.db.QueryRow(
		`INSERT INTO household_app.users (name, email, mobile, house)
		 VALUES ($1, $2, $3, $4)
		 returning id, name, email, mobile, house
		`,
		user.Name,
		user.Email,
		user.Mobile,
		user.House,
	).Scan(
		&createdUser.Id,
		&createdUser.Name,
		&createdUser.Email,
		&createdUser.Mobile,
		&createdUser.House,
	)
	if err != nil {
		return nil, err
	}
	log.Printf("User has been saved: %v", createdUser)
	return &createdUser, nil
}
