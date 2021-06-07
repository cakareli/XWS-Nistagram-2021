package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/service"
	"encoding/json"
	"fmt"
	"net/http"
)

type UserHandler struct {
	UserService *service.UserService
}

func(handler *UserHandler) Hello(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "Hello from controller!")
	handler.UserService.Hello()
}

func (handler *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var regularUserDto dto.RegularUserDTO
	err := json.NewDecoder(r.Body).Decode(&regularUserDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.CreateRegularUser(regularUserDto)
	if err != nil {
		if err.Error() == "Given username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) Update(w http.ResponseWriter, r *http.Request) {
	var userUpdateDto dto.UserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&userUpdateDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.UpdateRegularUser(userUpdateDto)
	if err != nil {
		if err.Error() == "Username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
}