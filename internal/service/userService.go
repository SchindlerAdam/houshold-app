package service

import (
	"houshold-app/internal/model"
	"houshold-app/internal/repository"
)

type UserService struct {
	repository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{userRepository}
}

func (userService *UserService) Save(user *model.CreateUser) (*model.User, error) {
	createdUser, err := userService.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}

func (userService *UserService) Update(id int, updateUser *model.UpdatedUser) (*model.User, error) {
	updatedUser, err := userService.repository.Update(id, updateUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}
