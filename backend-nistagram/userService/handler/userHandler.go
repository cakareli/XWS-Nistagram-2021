package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var regularUserRegistrationDto dto.RegularUserRegistration
	err := json.NewDecoder(r.Body).Decode(&regularUserRegistrationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.CreateRegularUser(regularUserRegistrationDto)
	if err != nil {
		if err.Error() == "username is already taken" {
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
	var userUpdateDto dto.RegularUserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&userUpdateDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.UserService.UpdateRegularUser(userUpdateDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (handler *UserHandler) FindUserById(w http.ResponseWriter, r *http.Request){


	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	regularUser, err := handler.UserService.FindUserById(id)
	/*
	collection := client.Database("thepolyglotdeveloper").Collection("people")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	err := collection.FindOne(ctx, Person{ID: id}).Decode(&person)*/
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(w).Encode(regularUser)

	/*
	keys, ok := r.URL.Query()["key"]

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		return
	}

	// Query()["key"] will return an array of items,
	key := keys[0]


	log.Println("Url Param 'key' is: " + string(key))
	user := handler.UserService.FindUserById(key)
	//user := handler.UserService.FindUserById("60c6838035a380ead199ffec")
	responseJSON, err2 := json.Marshal(user)
	if err2 != nil{
		fmt.Println(err2)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(responseJSON)
	fmt.Printf("JSON: \n", responseJSON)
	w.Header().Set("Content-Type", "application/json")*/
}