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

func (userService *UserService) Save(user *model.CreateUser) (*model.UserResponse, error) {
	createdUser, err := userService.repository.Create(user)
	if err != nil {
		return nil, err
	}
	return createdUser, nil

}

func (userService *UserService) Update(id int, updateUser *model.UpdateUser) (*model.UserResponse, error) {
	updatedUser, err := userService.repository.Update(id, updateUser)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (userService *UserService) GetAll() (*[]model.UserResponse, error) {
	users, err := userService.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (userService *UserService) Delete(id int) error {
	err := userService.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
