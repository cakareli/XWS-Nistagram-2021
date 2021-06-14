package service

import (
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/dto"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/model"
	"XWS-Nistagram-2021/backend-nistagram/authenticationService/repository"
	"fmt"
	//"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthenticationService struct {
	AuthenticationRepository *repository.AuthenticationRepository
}

func (service *AuthenticationService) Hello () {
	fmt.Printf("Hello from service!")
	service.AuthenticationRepository.Hello();
}

func (service *AuthenticationService) RegisterUser (dto dto.RegularUserRegistrationDTO) error {
	user := model.User{ Email: dto.Email, UserRole: 0, UserId: dto.UserId, Name: dto.Name, Surname: dto.Surname, Password: dto.Password, Username: dto.Username}
	err := service.AuthenticationRepository.RegisterUser(&user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthenticationService) UpdateUser (dto dto.RegularUserUpdateDTO) error {
	user, err := service.AuthenticationRepository.FindUserByUserId(dto.UserId)
	if err != nil {
		return err
	}
	user.Name = dto.Name;
	user.Surname = dto.Surname;
	user.Username = dto.Username;
	user.Email = dto.Email;
	err = service.AuthenticationRepository.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (service *AuthenticationService) FindByUsername (dto dto.LoginDTO) (*model.User, error){
	user, err := service.AuthenticationRepository.FindUserByUsername(dto.Username)
	if err != nil {
		return nil, err
	}
	return user, nil
}