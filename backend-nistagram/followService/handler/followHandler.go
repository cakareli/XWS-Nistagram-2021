package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/dto"
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

type FollowHandler struct {
	FollowService *service.FollowService
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

func (handler *FollowHandler) MuteFollowing(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followingId := params["followingId"]
	userIsMuted := handler.FollowService.MuteFollowing(loggedUserId, followingId)
	if !userIsMuted {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) AddToCloseFollowers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followingId := params["followingId"]
	userIsInClose := handler.FollowService.AddToCloseFollowers(loggedUserId, followingId)
	if !userIsInClose {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) BlockUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	userId := params["userId"]
	userIsBlocked := handler.FollowService.BlockUser(loggedUserId, userId)
	if !userIsBlocked {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func (handler *FollowHandler) UnblockUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	userId := params["userId"]
	userIsUnblocked := handler.FollowService.UnblockUser(loggedUserId, userId)
	if !userIsUnblocked {
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

func (handler *FollowHandler) FindAllUserFollowers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followers, err := handler.FollowService.FindAllUserFollowers(loggedUserId)
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

func (handler *FollowHandler) FindAllUserFollowings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followings, err := handler.FollowService.FindAllUserFollowings(loggedUserId)
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

func (handler *FollowHandler) FindAllUserBlockedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	blockedUsers, err := handler.FollowService.FindAllUserBlockedUsers(loggedUserId)
	if err != nil {
		if err.Error() == "no blocked users found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(blockedUsers)
}

func (handler *FollowHandler) FindAllUserMutedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	mutedUsers, err := handler.FollowService.FindAllUserMutedUsers(loggedUserId)
	if err != nil {
		if err.Error() == "no muted users found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(mutedUsers)
}

func (handler *FollowHandler) FindAllUserCloseFollowers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	closeFollowers, err := handler.FollowService.FindAllUserCloseFollowers(loggedUserId)
	if err != nil {
		if err.Error() == "no close followers found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(closeFollowers)
}

func (handler *FollowHandler) FindAllUserFollowRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	followRequests, err := handler.FollowService.FindAllUserFollowRequests(loggedUserId)
	if err != nil {
		if err.Error() == "no follow requests found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(followRequests)
}

func (handler *FollowHandler) FindAllFeedUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	loggedUserId := params["loggedUserId"]
	postsForFeed, err := handler.FollowService.FindAllPostsForFeed(loggedUserId)
	if err != nil {
		if err.Error() == "no feed users found" {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode("")
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(postsForFeed)
}