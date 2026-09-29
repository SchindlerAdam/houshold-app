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

func (repository UserRepository) Create(user *model.CreateUser) (*model.UserResponse, error) {
	var createdUser model.UserResponse
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

func (repository UserRepository) Update(id int, updateUser *model.UpdateUser) (*model.UserResponse, error) {
	userById, err := repository.GetById(id)

	if updateUser.Name != nil && *updateUser.Name != "" {
		userById.Name = *updateUser.Name
	}
	if updateUser.Email != nil && *updateUser.Email != "" {
		userById.Email = *updateUser.Email
	}
	if updateUser.Mobile != nil && *updateUser.Mobile != "" {
		userById.Mobile = *updateUser.Mobile
	}
	if updateUser.House != nil {
		userById.House = *updateUser.House
	}
	err = repository.db.QueryRow(
		`UPDATE household_app.users
		 SET name = $1, email = $2, mobile = $3, house = $4
		 WHERE id = $5
		 returning id, name, email, mobile, house
		`,
		userById.Name,
		userById.Email,
		userById.Mobile,
		userById.House,
		id,
	).Scan(
		&userById.Id,
		&userById.Name,
		&userById.Email,
		&userById.Mobile,
		&userById.House,
	)
	if err != nil {
		return nil, err
	}
	log.Printf("User has been updated: %v", userById)
	return userById, nil

}

func (repository UserRepository) GetById(id int) (*model.UserResponse, error) {
	var user model.UserResponse
	err := repository.db.QueryRow(
		`SELECT id, name, email, mobile, house FROM household_app.users WHERE id = $1`,
		id,
	).Scan(
		&user.Id,
		&user.Name,
		&user.Email,
		&user.Mobile,
		&user.House,
	)
	if err != nil {
		return nil, err
	}
	log.Printf("User has been retrieved: %v", user)
	return &user, nil
}

func (repository UserRepository) GetAll() (*[]model.UserResponse, error) {
	rows, err := repository.db.Query(
		`SELECT id, name, email, mobile, house FROM household_app.users WHERE is_deleted = FALSE`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []model.UserResponse
	for rows.Next() {
		var user model.UserResponse
		err := rows.Scan(
			&user.Id,
			&user.Name,
			&user.Email,
			&user.Mobile,
			&user.House,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &users, nil
}

func (repository UserRepository) Delete(id int) error {
	user, err := repository.GetById(id)
	if err != nil {
		return err
	}

	_, err = repository.db.Exec(
		`UPDATE household_app.users SET is_deleted = TRUE WHERE id = $1`,
		user.Id,
	)
	if err != nil {
		return err
	}
	log.Printf("User has been deleted: %v", user)
	return nil
}
