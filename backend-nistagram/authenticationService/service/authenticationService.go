package service

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/dto"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/repository"
	"fmt"
)

type AuthenticationService struct {
	AuthenticationRepository *repository.AuthenticationRepository
}

func (service *AuthenticationService) Hello () {
	fmt.Printf("Hello from service!")
	service.AuthenticationRepository.Hello();
}

func (service *AuthenticationService) RegisterUser (dto dto.RegularUserRegistrationDTO) error {
	user := model.User{Id: dto.Id, Email: dto.Email, UserRole: dto.Role, Name: dto.Name, Surname: dto.Surname, Password: dto.Password, Username: dto.Username}
	err := service.AuthenticationRepository.RegisterUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthenticationService) Login (dto dto.LoginDTO) (*model.User, error){
	user, err := service.AuthenticationRepository.FindUserByUsername(dto.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}