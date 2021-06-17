package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/dto"
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
	"encoding/json"
	"fmt"
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