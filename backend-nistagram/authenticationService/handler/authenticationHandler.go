package handler

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/service"
	"fmt"
	"net/http"
)

type AuthenticationHandler struct {
	AuthenticationService *service.AuthenticationService
}

func(handler *AuthenticationHandler) Hello(res http.ResponseWriter, req *http.Request){
		fmt.Fprint(res, "Hello from controller!")
		handler.AuthenticationService.Hello()
}