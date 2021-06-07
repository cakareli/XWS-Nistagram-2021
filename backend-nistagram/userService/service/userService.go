package service

import (
	"XWS-Nistagram-2021/backend-nistagram/userService/dto"
	"XWS-Nistagram-2021/backend-nistagram/userService/model"
	"XWS-Nistagram-2021/backend-nistagram/userService/repository"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func (service *UserService) Hello () {
	fmt.Printf("Hello from service!")
	service.UserRepository.Hello();
}

func (service *UserService) CreateRegularUser(regularUserDto dto.RegularUserDTO) error {
	fmt.Println("Creating regular user")

	if service.UserRepository.ExistByUsername(regularUserDto.Username) {
		return fmt.Errorf("Username is already taken")
	}

	var regularUser = createRegularUserFromRegularUserDTO(&regularUserDto)
	err := service.UserRepository.CreateRegularUser(regularUser)
	if err != nil {
		return err
	}
	return nil
}

func (service *UserService) UpdateRegularUser(userUpdateDto dto.UserUpdateDTO) error {
	fmt.Println("Updating regular user")

	if service.UserRepository.ExistByUsername(userUpdateDto.Username) {
		id, _ := primitive.ObjectIDFromHex(userUpdateDto.Id)
		if service.UserRepository.UsernameChanged(userUpdateDto.Username, id) {
			return fmt.Errorf("Username is already taken")
		}
	}

	var regularUser = createRegularUserFromUserUpdateDTO(&userUpdateDto)
	err := service.UserRepository.UpdateRegularUser(regularUser)
	if err != nil {
		return err
	}
	return nil
}

func createRegularUserFromRegularUserDTO(regularUserDto *dto.RegularUserDTO) *model.RegularUser{
	profilePrivacy := model.ProfilePrivacy{
		PrivacyType: regularUserDto.PrivacyType,
		AllMessageRequests: regularUserDto.AllMessageRequests,
		TagsAllowed: regularUserDto.TagsAllowed,
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
	regularUser.IsDisabled = regularUserDto.IsDisabled
	regularUser.UserRole = regularUserDto.UserRole
	regularUser.UserType = regularUserDto.UserType
	regularUser.Gender = regularUserDto.Gender

	return &regularUser
}

func createRegularUserFromUserUpdateDTO(userUpdateDto *dto.UserUpdateDTO) *model.RegularUser{
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
