package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/dto"
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowHandler struct {
	FollowService *service.FollowService
}

func(handler *FollowHandler) Hello(res http.ResponseWriter, req *http.Request){
	fmt.Fprint(res, "Hello from controller!")
	handler.FollowService.Hello()
}

func (handler *FollowHandler) FollowUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newFollowDto dto.NewFollowDTO
	err := json.NewDecoder(r.Body).Decode(&newFollowDto)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	userIsFollowed := handler.FollowService.FollowUser(newFollowDto)
	if !userIsFollowed {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) AcceptFollowRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followerId := params["followerId"]
	userIsAccepted := handler.FollowService.AcceptFollowRequest(loggedUserId, followerId)
	if !userIsAccepted {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) RemoveFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followingId := params["followingId"]
	followIsRemoved := handler.FollowService.RemoveFollowing(loggedUserId, followingId)
	if !followIsRemoved {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) RemoveFollower(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followerId := params["followerId"]
	followIsRemoved := handler.FollowService.RemoveFollower(loggedUserId, followerId)
	if !followIsRemoved {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) FindAllFollowers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followers, err := handler.FollowService.FindAllFollowers(loggedUserId)
	if err != nil {
		if err.Error() == "no followers found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(followers)
}

func (handler *FollowHandler) FindAllFollowings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followings, err := handler.FollowService.FindAllFollowings(loggedUserId)
	if err != nil {
		if err.Error() == "no followings found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(followings)
}