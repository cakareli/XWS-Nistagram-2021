package service

import (
	"fmt"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/repository"
)

type AuthenticationService struct {
	AuthenticationRepository *repository.AuthenticationRepository
}

func (service *AuthenticationService) Hello () {
	fmt.Printf("Hello from service!")
	service.AuthenticationRepository.Hello();
}
