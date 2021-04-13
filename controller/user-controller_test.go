package controller

import (
	"encoding/json"
	"github.com/WalletService/mocks"
	"github.com/WalletService/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	userService := new(mocks.IUserService)
	expectedResult := model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
	}
	userService.On("GetUserService", 1).Return(&expectedResult, nil)
	userController := NewUserController(userService)

	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/user/1", nil)
	w := httptest.NewRecorder()

	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.GetUser)
	r.ServeHTTP(w, req)

	actualResult := model.User{}
	json.NewDecoder(w.Body).Decode(&actualResult)
	assert.Equal(t, expectedResult,actualResult)
}
