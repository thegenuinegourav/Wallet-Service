package service

import (
	"errors"
	"github.com/WalletService/mocks"
	"github.com/WalletService/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetUserService(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	userService := NewUserService(userRepository)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
	}
	userRepository.On("GetUserById", 1).Return(&expectedResult1, nil)
	actualResult1, _ := userService.GetUserService(1)
	assert.Equal(t, &expectedResult1, actualResult1)

	t.Logf("Running Negative Case with error..........")
	userRepository.On("GetUserById", 2).Return(nil, errors.New("some error"))
	_, err := userService.GetUserService(2)
	assert.NotNil(t, err)
}

func TestGetUsersService(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	userService := NewUserService(userRepository)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := []model.User{
		model.User{
			Name:   "test_user",
			Email:  "test_email",
			Mobile: "test_mobile",
		},
	}
	userRepository.On("GetAllUsers").Return(&expectedResult1, nil).Once()
	actualResult1, _ := userService.GetUsersService()
	assert.Equal(t, &expectedResult1, actualResult1)

	t.Logf("Running Negative Case with error..........")
	userRepository.On("GetAllUsers").Return(nil, errors.New("some error")).Once()
	_, err := userService.GetUsersService()
	assert.NotNil(t, err)
}

func TestPostUserService(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	userService := NewUserService(userRepository)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name:   "test_user",
		Email:  "test_email",
		Mobile: "test_mobile",
	}
	userRepository.On("CreateUser", &expectedResult1).Return(&expectedResult1, nil)
	actualResult1, _ := userService.PostUserService(&expectedResult1)
	assert.Equal(t, &expectedResult1, actualResult1)

	t.Logf("Running Negative Case with error..........")
	invalidUser := model.User{
		Name:   "test_user",
	}
	userRepository.On("CreateUser", &invalidUser).Return(nil, errors.New("some error"))
	_, err := userService.PostUserService(&invalidUser)
	assert.NotNil(t, err)
}

func TestUpdateUserService(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	userService := NewUserService(userRepository)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name:   "test_user",
		Email:  "test_email",
		Mobile: "test_mobile",
	}
	userRepository.On("GetUserById",1).Return(&expectedResult1, nil)
	userRepository.On("UpdateUser",&expectedResult1).Return(&expectedResult1, nil)
	actualResult1, _ := userService.UpdateUserService(1, &expectedResult1)
	assert.Equal(t, &expectedResult1, actualResult1)


	t.Logf("Running Negative Case with Not Found..........")
	userRepository.On("GetUserById",2).Return(nil, errors.New("Not Found"))
	_, err := userService.UpdateUserService(2, &expectedResult1)
	assert.NotNil(t, err)

	t.Logf("Running Negative Case with error..........")
	invalidUser := model.User{
		Name:   "test_user",
	}
	userRepository.On("GetUserById",3).Return(&invalidUser, nil)
	userRepository.On("UpdateUser",&invalidUser).Return(nil, errors.New("some error"))
	_, err = userService.UpdateUserService(3, &invalidUser)
	assert.NotNil(t, err)
}

func TestDeleteUserService(t *testing.T) {
	userRepository := new(mocks.IUserRepository)
	userService := NewUserService(userRepository)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name:   "test_user",
		Email:  "test_email",
		Mobile: "test_mobile",
	}
	userRepository.On("GetUserById",1).Return(&expectedResult1, nil)
	userRepository.On("DeleteUser",&expectedResult1).Return(nil)
	err := userService.DeleteUserService(1)
	assert.Nil(t, err)

	t.Logf("Running Negative Case with Not Found..........")
	userRepository.On("GetUserById",2).Return(nil, errors.New("Not Found"))
	err = userService.DeleteUserService(2)
	assert.NotNil(t, err)

	t.Logf("Running Negative Case with error..........")
	invalidUser := model.User{
		Name:   "test_user",
	}
	userRepository.On("GetUserById",3).Return(&invalidUser, nil)
	userRepository.On("DeleteUser",&invalidUser).Return(errors.New("some error"))
	err = userService.DeleteUserService(3)
	assert.NotNil(t, err)
}



