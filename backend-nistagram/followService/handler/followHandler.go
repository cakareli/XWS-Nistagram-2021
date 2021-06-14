package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/followService/service"
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