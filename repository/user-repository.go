package repository

import (
	"github.com/jinzhu/gorm"
	. "github.com/WalletService/model"
)

type IUserRepository interface{
	GetUserById(id int) (*User, error)
	GetAllUsers() (*[]User, error)
	CreateUser(user *User) (*User, error)
	UpdateUser(user *User) (*User, error)
	DeleteUser(user *User) error
}

type userRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (userRepository *userRepository) GetUserById(id int) (*User, error) {
	var user User
	result := userRepository.DB.First(&user,id)
	return &user, result.Error
}

func (userRepository *userRepository) GetAllUsers() (*[]User, error) {
	var user []User
	result := userRepository.DB.Find(&user)
	return &user, result.Error
}

func (userRepository *userRepository) CreateUser(user *User) (*User, error) {
	result := userRepository.DB.Create(user)
	return user, result.Error
}

func (userRepository *userRepository) UpdateUser(user *User) (*User, error) {
	result := userRepository.DB.Save(user)
	return user, result.Error
}

func (userRepository *userRepository) DeleteUser(user *User) error {
	result := userRepository.DB.Delete(user)
	return result.Error
}
