package service

import (
	"github.com/WalletService/repository"
	. "github.com/WalletService/model"
)

type IUserService interface{
	GetUserService(id int) (*User, error)
	GetUsersService() (*[]User, error)
	PostUserService(user *User) (*User, error)
	UpdateUserService(id int, user *User) (*User, error)
	DeleteUserService(id int) error
}

type userService struct {}

var (
	userRepository repository.IUserRepository
)

func NewUserService(repository repository.IUserRepository) IUserService{
	userRepository = repository
	return &userService{}
}

func (userService *userService) GetUserService(id int) (*User, error) {
	return userRepository.GetUserById(id)
}

func (userService *userService) GetUsersService() (*[]User, error) {
	return userRepository.GetAllUsers()
}

func (userService *userService) PostUserService(user *User) (*User, error) {
	return userRepository.CreateUser(user)
}

func (userService *userService) UpdateUserService(id int, user *User) (*User, error) {
	res, err := userRepository.GetUserById(id)
	if err != nil {
		return nil, err
	}
	updateUserGivenFields(res, user)
	return userRepository.UpdateUser(res)
}

func (userService *userService) DeleteUserService(id int) error {
	res, err := userRepository.GetUserById(id)
	if err != nil {
		return err
	}
	return userRepository.DeleteUser(res)
}

func updateUserGivenFields(u1 *User, u2 *User) {
	if len(u2.Name)!=0 {
		u1.Name = u2.Name
	}
	if len(u2.Email)!=0 {
		u1.Email = u2.Email
	}
	if len(u2.Mobile)!=0 {
		u1.Mobile = u2.Mobile
	}
}
