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
	"os"
)

type RegularUserService struct {
	RegularUserRepository *repository.RegularUserRepository
}

func (service *RegularUserService) Register(regularUserRegistrationDto dto.RegularUserRegistrationDTO) error {
	fmt.Println("Creating regular user")

	if service.RegularUserRepository.ExistByUsername(regularUserRegistrationDto.Username) {
		return fmt.Errorf("username is already taken")
	}

	var regularUser = createRegularUserFromRegularUserRegistrationDTO(&regularUserRegistrationDto)
	createdUserId, err := service.RegularUserRepository.Register(regularUser)
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
	requestUrl := fmt.Sprintf("http://%s:%s/register", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) Update(regularUserUpdateDto dto.RegularUserUpdateDTO) error {
	fmt.Println("Updating regular user")

	if service.RegularUserRepository.ExistByUsername(regularUserUpdateDto.Username) {
		id, _ := primitive.ObjectIDFromHex(regularUserUpdateDto.Id)
		if service.RegularUserRepository.UsernameChanged(regularUserUpdateDto.Username, id) {
			return fmt.Errorf("username is already taken")
		}
	}

	var regularUser = createRegularUserFromRegularUserUpdateDTO(&regularUserUpdateDto)
	err := service.RegularUserRepository.Update(regularUser)
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
	requestUrl := fmt.Sprintf("http://%s:%s/update", os.Getenv("AUTHENTICATION_SERVICE_DOMAIN"), os.Getenv("AUTHENTICATION_SERVICE_PORT"))
	resp, err := http.Post(requestUrl, "application/json", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println(resp.StatusCode)
	return nil
}

func (service *RegularUserService) FindUserById(userId primitive.ObjectID) (*model.RegularUser, error){
	fmt.Print("Searching for logged user...")
	regularUser, err := service.RegularUserRepository.FindUserById(userId)
	if err != nil {
		return nil, err
	}
	return regularUser, err
}

func (service *RegularUserService) FindUserByUsername(username string) (*dto.RegularUserPostDTO, error){
	regularUser, err := service.RegularUserRepository.FindUserByUsername(username)
	if err != nil {
		return nil, err
	}
	regularUserPostDto := createRegularUserPostDTOFromRegularUser(regularUser)
	return regularUserPostDto, nil
}

func createRegularUserPostDTOFromRegularUser(regularUser *model.RegularUser) *dto.RegularUserPostDTO {
	var regularUserPostDto dto.RegularUserPostDTO
	regularUserPostDto.Id = regularUser.Id.Hex()
	regularUserPostDto.PrivacyType = &regularUser.ProfilePrivacy.PrivacyType

	return &regularUserPostDto
}

func createRegularUserFromRegularUserRegistrationDTO(regularUserDto *dto.RegularUserRegistrationDTO) *model.RegularUser{
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
