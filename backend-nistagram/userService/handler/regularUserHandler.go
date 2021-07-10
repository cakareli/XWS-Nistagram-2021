package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type RegularUserHandler struct {
	RegularUserService *service.RegularUserService
}

func (handler *RegularUserHandler) Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var regularUserRegistrationDto dto.RegularUserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&regularUserRegistrationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.Register(regularUserRegistrationDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *RegularUserHandler) RegisterAgent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var regularUserRegistrationDto dto.RegularUserRegistrationDTO
	err := json.NewDecoder(r.Body).Decode(&regularUserRegistrationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.RegisterAgent(regularUserRegistrationDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (handler *RegularUserHandler) UpdatePersonalInformations(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var userUpdateDto dto.RegularUserUpdateDTO
	err := json.NewDecoder(r.Body).Decode(&userUpdateDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdatePersonalInformations(userUpdateDto)
	if err != nil {
		if err.Error() == "username is already taken" {
			w.WriteHeader(http.StatusConflict)
		} else {
			w.WriteHeader(http.StatusBadRequest)
		}
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *RegularUserHandler) UpdateProfilePrivacy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var profilePrivacyDto dto.ProfilePrivacyDTO
	err := json.NewDecoder(r.Body).Decode(&profilePrivacyDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdateProfilePrivacy(profilePrivacyDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *RegularUserHandler) DeleteRegularUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	var deleteUserDto dto.DeleteUserDTO
	err1 := json.NewDecoder(r.Body).Decode(&deleteUserDto)
	if err1 != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	err := handler.RegularUserService.DeleteRegularUser(deleteUserDto)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (handler *RegularUserHandler) CreateRegularUserPostDTOByUsername(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPostDto, err := handler.RegularUserService.CreateRegularUserPostDTOByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(regularUserPostDto)
}

func (handler *RegularUserHandler) FindRegularUserByUsername(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPostDto, err := handler.RegularUserService.FindRegularUserByUsername(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regularUserPostDto)
}

func (handler *RegularUserHandler) FindRegularUserLikedAndDislikedPosts(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	regularUserPostDto, err := handler.RegularUserService.FindRegularUserLikedAndDislikedPosts(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(regularUserPostDto)
}

func (handler *RegularUserHandler) FindUserById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	regularUser, err := handler.RegularUserService.FindUserById(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(regularUser)
}

func (handler *RegularUserHandler) GetAllPublicRegularUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	allRegularUsersDto, err := handler.RegularUserService.GetAllPublicRegularUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allRegularUsersDto)
	}
}

func (handler *RegularUserHandler) GetAllRegularUsers(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	allRegularUsersDto, err := handler.RegularUserService.GetAllPublicRegularUsers()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allRegularUsersDto)
	}
}

func (handler *RegularUserHandler) GetUserSearchResults(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	searchInput := param["searchInput"]
	searchUsers, err := handler.RegularUserService.GetUserSearchResults(searchInput)
	searchPostsJson, err := json.Marshal(searchUsers)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write(searchPostsJson)
	}
}

func (handler *RegularUserHandler) FindUsersByIds(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	var usersIds []string
	err := json.NewDecoder(r.Body).Decode(&usersIds)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userFollowDtos, err := handler.RegularUserService.FindUsersByIds(usersIds)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(userFollowDtos)
}

func (handler *RegularUserHandler) UpdateLikedPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postLikeDTO dto.UpdatePostLikeAndDislikeDTO
	err := json.NewDecoder(r.Body).Decode(&postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdateLikedPosts(postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (handler *RegularUserHandler) UpdateDislikedPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postLikeDTO dto.UpdatePostLikeAndDislikeDTO
	err := json.NewDecoder(r.Body).Decode(&postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.UpdateDislikedPosts(postLikeDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (handler *RegularUserHandler) SavePost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var postSaveDTO dto.PostSaveDTO
	err := json.NewDecoder(r.Body).Decode(&postSaveDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.SavePost(postSaveDTO)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
	w.Header().Set("Content-Type", "application/json")
}

func (handler *RegularUserHandler) GetAllSavedPostsByUsername(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	param := mux.Vars(r)
	username := param["username"]
	savedPosts, err := handler.RegularUserService.FindRegularUserSavedPosts(username)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedPosts)
}

func (handler *RegularUserHandler) VerifyUser(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var userVerificationDto dto.UserVerificationDTO
	err := json.NewDecoder(r.Body).Decode(&userVerificationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	err = handler.RegularUserService.VerifyUser(userVerificationDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}
/*
func (handler *RegularUserHandler) GetAllAgentRequests(w http.ResponseWriter, r *http.Request){
	w.Header().Set("content-type", "application/json")
	allRegularUsersDto, err := handler.RegularUserService.GetAllAgentRequests()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(allRegularUsersDto)
	}
}*/