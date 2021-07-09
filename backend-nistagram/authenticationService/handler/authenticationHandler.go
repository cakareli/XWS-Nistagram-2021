package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/dto"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/service"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/util"
	"encoding/json"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

type AuthenticationHandler struct {
	AuthenticationService *service.AuthenticationService
}

func (handler *AuthenticationHandler) RegisterUser (res http.ResponseWriter, req *http.Request) {
	var regularUserRegistrationDTO dto.RegularUserRegistrationDTO
	err := json.NewDecoder(req.Body).Decode(&regularUserRegistrationDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return;
	}

	err = handler.AuthenticationService.RegisterUser(regularUserRegistrationDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return;
	}
	res.WriteHeader(http.StatusCreated);
}

func (handler *AuthenticationHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	id := param["id"]
	err := handler.AuthenticationService.DeleteUser(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *AuthenticationHandler) UpdateUser (res http.ResponseWriter, req *http.Request) {
	var regularUserUpdateDTO dto.RegularUserUpdateDTO
	err := json.NewDecoder(req.Body).Decode(&regularUserUpdateDTO)
	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		return;
	}
	fmt.Println(err)
	err = handler.AuthenticationService.UpdateUser(regularUserUpdateDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return;
	}
	res.WriteHeader(http.StatusOK);
}

func(handler *AuthenticationHandler) Login(res http.ResponseWriter, req *http.Request){
	var loginDTO dto.LoginDTO
	err := json.NewDecoder(req.Body).Decode(&loginDTO)
	if err !=nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	var user *model.User
	user, err = handler.AuthenticationService.FindByUsername(loginDTO)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	if(user.Password != loginDTO.Password){
		res.WriteHeader(http.StatusUnauthorized)
		return
	}
	token, err := util.CreateJWT(user.UserId, &user.UserRole, user.Username)
	response := dto.LoginResponseDTO{
		Token: token,
	}
	responseJSON, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(responseJSON)
	res.Header().Set("Content-Type", "application/json")
}