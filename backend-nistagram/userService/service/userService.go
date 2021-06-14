package service

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"XWS-Nistagram-2021/backend-nistagram/userService/repository"
	"bytes"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) Hello () {
	fmt.Printf("Hello from service!")
	service.UserRepository.Hello()
}

func (service *UserService) CreateRegularUser(regularUserRegistrationDto dto.RegularUserRegistration) error {
	fmt.Println("Creating regular user")

	if service.UserRepository.ExistByUsername(regularUserRegistrationDto.Username) {
		return fmt.Errorf("username is already taken")
	}

	var regularUser = createRegularUserFromRegularUserRegistrationDTO(&regularUserRegistrationDto)
	createdUserId, err := service.UserRepository.CreateRegularUser(regularUser)
	if err != nil {
		return err
	}
	postBody, _ := json.Marshal(map[string]string{
		"userId": createdUserId,
		"email": regularUserRegistrationDto.Email,
		"password": regularUserRegistrationDto.Password,
		"username": regularUserRegistrationDto.Username,
		"name": regularUserRegistrationDto.Name,
		"surname": regularUserRegistrationDto.Surname,
	})
	resp, err := http.Post("http://authentication-service:8081/register", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *UserService) UpdateRegularUser(regularUserUpdateDto dto.RegularUserUpdateDTO) error {
	fmt.Println("Updating regular user")

	if service.UserRepository.ExistByUsername(regularUserUpdateDto.Username) {
		id, _ := primitive.ObjectIDFromHex(regularUserUpdateDto.Id)
		if service.UserRepository.UsernameChanged(regularUserUpdateDto.Username, id) {
			return fmt.Errorf("username is already taken")
		}
	}

	var regularUser = createRegularUserFromRegularUserUpdateDTO(&regularUserUpdateDto)
	err := service.UserRepository.UpdateRegularUser(regularUser)
	if err != nil {
		return err
	}

	postBody, _ := json.Marshal(map[string]string{
		"_id": regularUserUpdateDto.Id,
		"email": regularUserUpdateDto.Email,
		"username": regularUserUpdateDto.Username,
		"name": regularUserUpdateDto.Name,
		"surname": regularUserUpdateDto.Surname,
	})

	resp, err := http.Post("http://authentication-service:8081/update", "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *UserService) FindUserById(userId primitive.ObjectID) (*model.RegularUser, error){
	fmt.Print("Searching for logged user...")
	regularUser, err := service.UserRepository.FindUserById(userId)
	return regularUser, err
}

func createRegularUserFromRegularUserRegistrationDTO(regularUserDto *dto.RegularUserRegistration) *model.RegularUser{
	profilePrivacy := model.ProfilePrivacy{
		PrivacyType: model.PrivacyType(0),
		AllMessageRequests: true,
		TagsAllowed: true,
	}
	var regularUser model.RegularUser
	regularUser.Name = regularUserDto.Name
	regularUser.Surname = regularUserDto.Surname
	regularUser.Username = regularUserDto.Username
	regularUser.Password = regularUserDto.Password
	regularUser.Email = regularUserDto.Email
	regularUser.PhoneNumber = regularUserDto.PhoneNumber
	regularUser.BirthDate = regularUserDto.BirthDate
	regularUser.Biography = regularUserDto.Biography
	regularUser.WebSite = regularUserDto.WebSite
	regularUser.ProfilePrivacy = profilePrivacy
	regularUser.IsDisabled = false
	regularUser.UserRole = model.UserRole(0)
	regularUser.UserType = model.UserType(0)
	regularUser.Gender = regularUserDto.Gender

	return &regularUser
}

func createRegularUserFromRegularUserUpdateDTO(userUpdateDto *dto.RegularUserUpdateDTO) *model.RegularUser{
	id, _ := primitive.ObjectIDFromHex(userUpdateDto.Id)
	var regularUser model.RegularUser
	regularUser.Id = id
	regularUser.Name = userUpdateDto.Name
	regularUser.Surname = userUpdateDto.Surname
	regularUser.Username = userUpdateDto.Username
	regularUser.Email = userUpdateDto.Email
	regularUser.PhoneNumber = userUpdateDto.PhoneNumber
	regularUser.Gender= userUpdateDto.Gender
	regularUser.BirthDate = userUpdateDto.BirthDate
	regularUser.Biography = userUpdateDto.Biography
	regularUser.WebSite = userUpdateDto.WebSite

	return &regularUser
}


