package controller

import (
	"database/sql"
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

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
	}
	userService.On("GetUserService", 1).Return(&expectedResult1, nil)
	userController := NewUserController(userService)
	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/user/1", nil)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.GetUser)
	r.ServeHTTP(w, req)
	actualResult1 := model.User{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, expectedResult1,actualResult1)

	t.Logf("Running Negative Case with error of Not Found...........")
	expectedResult2 := map[string]string{
		"error": "User not found",
	}
	userService.On("GetUserService", 2).Return(nil, sql.ErrNoRows)
	//userController := NewUserController(userService)
	req = httptest.NewRequest("GET", "http://localhost:8080/api/v1/user/2", nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.GetUser)
	r.ServeHTTP(w, req)
	actualResult2 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult2)
	assert.Equal(t, expectedResult2,actualResult2)

	t.Logf("Running Negative Case with error of Invalid Param...........")
	expectedResult3 := map[string]string{
		"error": "Invalid user ID",
	}
	req = httptest.NewRequest("GET", "http://localhost:8080/api/v1/user/1abs", nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.GetUser)
	r.ServeHTTP(w, req)
	actualResult3 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult3)
	assert.Equal(t, expectedResult3,actualResult3)
}


