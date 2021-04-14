package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"github.com/WalletService/mocks"
	"github.com/WalletService/model"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http/httptest"
	"testing"
)

func TestGetUser(t *testing.T) {
	userService := new(mocks.IUserService)
	userController := NewUserController(userService)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
	}
	userService.On("GetUserService", 1).Return(&expectedResult1, nil)
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

func TestGetUsers(t *testing.T) {
	userService := new(mocks.IUserService)
	userController := NewUserController(userService)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := []model.User{
		model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
		},
	}
	userService.On("GetUsersService").Return(&expectedResult1, nil).Once()
	req := httptest.NewRequest("GET", "http://localhost:8080/api/v1/user", nil)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user", userController.GetUsers)
	r.ServeHTTP(w, req)
	actualResult1 := []model.User{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, expectedResult1,actualResult1)

	t.Logf("Running Negative Case with error of Internal Server Error...........")
	expectedResult2 := map[string]string{"error":"something went wrong"}
	userService.On("GetUsersService").Return(nil, errors.New("something went wrong")).Once()
	req = httptest.NewRequest("GET", "http://localhost:8080/api/v1/user", nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user", userController.GetUsers)
	r.ServeHTTP(w, req)
	actualResult2 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult2)
	assert.Equal(t, expectedResult2,actualResult2)
}

func TestPostUser(t *testing.T) {
	userService := new(mocks.IUserService)
	userController := NewUserController(userService)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
	}
	userService.On("PostUserService", &expectedResult1).Return(&expectedResult1, nil)
	body,_ := json.Marshal(&expectedResult1)
	req := httptest.NewRequest("POST", "http://localhost:8080/api/v1/user", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user", userController.PostUser)
	r.ServeHTTP(w, req)
	actualResult1 := model.User{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, expectedResult1,actualResult1)

	t.Logf("Running Negative Case with error of Invalid Payload...........")
	expectedResult2 := map[string]string{
		"error": "Invalid request payload",
	}
	req = httptest.NewRequest("POST", "http://localhost:8080/api/v1/user", bytes.NewBuffer([]byte(`invalid json`)))
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user", userController.PostUser)
	r.ServeHTTP(w, req)
	actualResult2 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult2)
	assert.Equal(t, expectedResult2,actualResult2)
}

func TestPutUser(t *testing.T) {
	userService := new(mocks.IUserService)
	userController := NewUserController(userService)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := model.User{
		Name: "test_user",
		Email: "test_email",
		Mobile: "test_mobile",
	}
	userService.On("UpdateUserService", 1, &expectedResult1).Return(&expectedResult1, nil)
	body,_ := json.Marshal(&expectedResult1)
	req := httptest.NewRequest("PUT", "http://localhost:8080/api/v1/user/1", bytes.NewBuffer(body))
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.PutUser)
	r.ServeHTTP(w, req)
	actualResult1 := model.User{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, expectedResult1,actualResult1)

	t.Logf("Running Negative Case with error of Invalid Payload...........")
	expectedResult2 := map[string]string{
		"error": "Invalid request payload",
	}
	req = httptest.NewRequest("PUT", "http://localhost:8080/api/v1/user/1", bytes.NewBuffer([]byte(`invalid json`)))
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.PutUser)
	r.ServeHTTP(w, req)
	actualResult2 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult2)
	assert.Equal(t, expectedResult2,actualResult2)

	t.Logf("Running Negative Case with error of Invalid Param...........")
	expectedResult3 := map[string]string{
		"error": "Invalid user ID",
	}
	req = httptest.NewRequest("PUT", "http://localhost:8080/api/v1/user/1abs", nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.PutUser)
	r.ServeHTTP(w, req)
	actualResult3 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult3)
	assert.Equal(t, expectedResult3,actualResult3)

	t.Logf("Running Negative Case with error of Not Found...........")
	expectedResult4 := map[string]string{
		"error": "User not found",
	}
	userService.On("UpdateUserService", 2, &expectedResult1).Return(nil, sql.ErrNoRows)
	req = httptest.NewRequest("PUT", "http://localhost:8080/api/v1/user/2", bytes.NewBuffer(body))
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.PutUser)
	r.ServeHTTP(w, req)
	actualResult4 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult4)
	assert.Equal(t, expectedResult4,actualResult4)


}

func TestDeleteUser(t *testing.T) {
	userService := new(mocks.IUserService)
	userController := NewUserController(userService)

	t.Logf("Running Positive Case with valid output..........")
	expectedResult1 := map[string]string{
		"result" : "success",
	}
	userService.On("DeleteUserService", 1).Return(nil)
	req := httptest.NewRequest("DELETE", "http://localhost:8080/api/v1/user/1", nil)
	w := httptest.NewRecorder()
	r := mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.DeleteUser)
	r.ServeHTTP(w, req)
	actualResult1 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult1)
	assert.Equal(t, expectedResult1,actualResult1)

	t.Logf("Running Negative Case with error of Not Found...........")
	expectedResult2 := map[string]string{
		"error": "User not found",
	}
	userService.On("DeleteUserService", 2).Return(sql.ErrNoRows)
	req = httptest.NewRequest("DELETE", "http://localhost:8080/api/v1/user/2", nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.DeleteUser)
	r.ServeHTTP(w, req)
	actualResult2 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult2)
	assert.Equal(t, expectedResult2,actualResult2)

	t.Logf("Running Negative Case with error of Invalid Param...........")
	expectedResult3 := map[string]string{
		"error": "Invalid user ID",
	}
	req = httptest.NewRequest("DELETE", "http://localhost:8080/api/v1/user/1abs", nil)
	w = httptest.NewRecorder()
	r = mux.NewRouter()
	r.HandleFunc("/api/v1/user/{id}", userController.DeleteUser)
	r.ServeHTTP(w, req)
	actualResult3 := map[string]string{}
	json.NewDecoder(w.Body).Decode(&actualResult3)
	assert.Equal(t, expectedResult3,actualResult3)
}


